// main.go
package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	//go:embed slide_window.lua
	luaScript string
)

// getClientIP 从请求中尽量取到真实客户端 IP（简单版）
func getClientIP(r *http.Request) string {
	// X-Forwarded-For 优先
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	// X-Real-IP 其次
	if xr := r.Header.Get("X-Real-Ip"); xr != "" {
		return xr
	}
	// 最后用 RemoteAddr（去掉端口）
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func RateLimitMiddleware(rdb *redis.Client, limit int, window time.Duration) func(http.Handler) http.Handler {
	// 把脚本包装成 redis.Script，这样 client 会自动 ScriptLoad/EvalSha
	script := redis.NewScript(luaScript)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			clientID := getClientIP(r) // 这里按 IP 限流，生产可改为 API Key / UserID
			key := fmt.Sprintf("rate:sliding:%s", clientID)

			nowMs := time.Now().UnixMilli()
			// ARGV 顺序必须和你的 Lua 脚本一致： limit, window(ms), now(ms)
			res, err := script.Run(ctx, rdb, []string{key},
				limit,
				int(window.Milliseconds()),
				nowMs,
			).Result()
			if err != nil {
				// 脚本执行错误（例如 Redis 无响应）时的策略：这里允许通过并记录错误
				log.Printf("rate limit script error: %v", err)
				next.ServeHTTP(w, r)
				return
			}

			allowed, ok := res.(int64)
			if !ok {
				// 安全兜底
				log.Printf("unexpected script result type: %T -> %#v", res, res)
				next.ServeHTTP(w, r)
				return
			}

			if allowed == 1 {
				next.ServeHTTP(w, r)
				return
			}
			// 被限流
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, `{"error":"rate limit exceeded"}`)
		})
	}
}

func main() {
	// Redis client - 修改为你自己的 Redis 地址/密码/DB
	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.154.97.245:6379",
		Password: "", // no password set
		DB:       0,
	})

	// 简单的 ping 验证连接
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("cannot connect to redis: %v", err)
	}

	// 例子：允许每个 client 5 次 / 10 秒（滑动窗口）
	limit := 5
	window := 10 * time.Second

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world")
	})

	handler := RateLimitMiddleware(rdb, limit, window)(mux)

	addr := ":8080"
	log.Printf("starting server on %s, limit=%d per %s", addr, limit, window)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
