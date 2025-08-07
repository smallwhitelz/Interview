package _2_concurrency_map

import (
	"fmt"
	"sync"
)

func Map() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				m[j] = id
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("写入完成，map长度：", len(m))
}
