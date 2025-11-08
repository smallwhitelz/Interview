package _3_sync

import (
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				// 这句话会打印几次？
				// 只会打印一次
				t.Log("ABC")
			})
		}()
	}
}
