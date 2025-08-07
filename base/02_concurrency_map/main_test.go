package _2_concurrency_map

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	v, ok := trySwap()
	fmt.Println(ok, v)
}

func trySwap() (*any, bool) {
	return nil, false
}
