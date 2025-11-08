package test

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	newLogger := logger.New(
		// 自定义日志输出（这里设置为标准输出）
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info, // 设置日志级别为 Info，打印 SQL
			Colorful:                  true,        // 使用彩色日志
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
		},
	)
	// 记得换成你的数据库配置
	dsn := "root:root@tcp(43.154.97.245:13316)/webook_intr?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}
