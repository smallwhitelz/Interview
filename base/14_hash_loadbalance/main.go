package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// ConsistentHash 实现（虚拟节点）
type ConsistentHash struct {
	replicas int               // 每台真实机器的虚拟节点数
	hashRing []uint32          // 已排序的 hash 值列表
	hashMap  map[uint32]string // hash -> node
	lock     sync.RWMutex
}

func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		replicas: replicas,
		hashMap:  make(map[uint32]string),
	}
}

func (c *ConsistentHash) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// Add 向环上添加一个真实节点（同时添加 replicas 个虚拟节点）
func (c *ConsistentHash) Add(node string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for i := 0; i < c.replicas; i++ {
		virtualNodeKey := node + "#" + strconv.Itoa(i)
		h := c.hashKey(virtualNodeKey)
		c.hashRing = append(c.hashRing, h)
		c.hashMap[h] = node
	}
	sort.Slice(c.hashRing, func(i, j int) bool { return c.hashRing[i] < c.hashRing[j] })
}

// Remove 从环上删除一个真实节点（以及其虚拟节点）
func (c *ConsistentHash) Remove(node string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	// 逐个删除该节点的虚拟节点
	for i := 0; i < c.replicas; i++ {
		virtualNodeKey := node + "#" + strconv.Itoa(i)
		h := c.hashKey(virtualNodeKey)
		delete(c.hashMap, h)
		// 从 hashRing 中移除 h
		idx := sort.Search(len(c.hashRing), func(i int) bool { return c.hashRing[i] >= h })
		if idx < len(c.hashRing) && c.hashRing[idx] == h {
			// 删除元素
			c.hashRing = append(c.hashRing[:idx], c.hashRing[idx+1:]...)
		}
	}
}

// Get 根据 key 获取对应的真实节点
func (c *ConsistentHash) Get(key string) (string, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if len(c.hashRing) == 0 {
		return "", fmt.Errorf("no nodes available")
	}
	h := c.hashKey(key)
	// 二分查找第一个 >= h 的位置
	idx := sort.Search(len(c.hashRing), func(i int) bool { return c.hashRing[i] >= h })
	if idx == len(c.hashRing) {
		// 如果超过最大值，环回到第一个
		idx = 0
	}
	node := c.hashMap[c.hashRing[idx]]
	return node, nil
}

func main() {
	c := NewConsistentHash(100) // 推荐虚拟节点数量依场景调试，常见 100-200
	c.Add("srv-A:8080")
	c.Add("srv-B:8080")
	c.Add("srv-C:8080")

	var keys = []string{
		// IPs (15)
		"192.0.2.1",
		"192.0.2.2",
		"192.0.2.3",
		"192.0.2.4",
		"192.0.2.5",
		"192.0.2.6",
		"192.0.2.7",
		"192.0.2.8",
		"192.0.2.9",
		"192.0.2.10",
		"198.51.100.1",
		"198.51.100.2",
		"198.51.100.3",
		"198.51.100.4",
		"198.51.100.5",

		// user ids (30)
		"user-001", "user-002", "user-003", "user-004", "user-005",
		"user-006", "user-007", "user-008", "user-009", "user-010",
		"user-011", "user-012", "user-013", "user-014", "user-015",
		"user-016", "user-017", "user-018", "user-019", "user-020",
		"user-021", "user-022", "user-023", "user-024", "user-025",
		"user-026", "user-027", "user-028", "user-029", "user-030",

		// sessions (20)
		"session-aaa-001", "session-aaa-002", "session-aaa-003", "session-aaa-004", "session-aaa-005",
		"session-aaa-006", "session-aaa-007", "session-aaa-008", "session-aaa-009", "session-aaa-010",
		"session-aaa-011", "session-aaa-012", "session-aaa-013", "session-aaa-014", "session-aaa-015",
		"session-aaa-016", "session-aaa-017", "session-aaa-018", "session-aaa-019", "session-aaa-020",

		// orders (20)
		"order-1001", "order-1002", "order-1003", "order-1004", "order-1005",
		"order-1006", "order-1007", "order-1008", "order-1009", "order-1010",
		"order-1011", "order-1012", "order-1013", "order-1014", "order-1015",
		"order-1016", "order-1017", "order-1018", "order-1019", "order-1020",

		// UUID-like keys (15)
		"3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"9b2d7a10-1f3c-4e6d-b8a6-e3f2a5b6c7d8",
		"4a7f9d24-6c3b-4f2d-8e9a-0c1b2d3e4f50",
		"6c1d2f33-8e4b-4a6d-b7c8-9f0a1b2c3d4e",
		"b7a1c2d3-e4f5-6a7b-8c9d-0e1f2a3b4c5d",
		"f1e2d3c4-b5a6-7890-1234-56789abcdef0",
		"a0b1c2d3-e4f5-6789-0abc-def123456789",
		"123e4567-e89b-12d3-a456-426614174000",
		"c56a4180-65aa-42ec-a945-5fd21dec0538",
		"7d9f1a22-33b4-4c6d-9e0f-1a2b3c4d5e6f",
		"8f14e45f-ceea-4f2b-9b8d-2c3a4b5c6d7e",
		"2b6f1e88-9a3c-42b7-8d3f-4a5b6c7d8e9f",
		"5f6e7d8c-9b0a-4c3d-b2e1-0f1a2b3c4d5e",
		"0a1b2c3d-4e5f-6789-abcd-ef0123456789",
		"de305d54-75b4-431b-adb2-eb6b9e546014",
	}
	for _, k := range keys {
		n, _ := c.Get(k)
		fmt.Printf("key=%s -> node=%s\n", k, n)
	}

	// 模拟移除一个节点，观察大多数 key 不变（只有部分 key 改变）
	fmt.Println("=== remove srv-B ===")
	c.Remove("srv-B:8080")
	for _, k := range keys {
		n, _ := c.Get(k)
		fmt.Printf("key=%s -> node=%s\n", k, n)
	}
}
