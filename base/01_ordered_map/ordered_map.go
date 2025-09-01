package _1_ordered_map

import "sync"

// 问题：如何在 Go 中按照特定顺序遍历 map,怎么做？
// 1. 在往map中插入数据前，维持一个key的切片，保证map的顺序，而后在遍历的时候就遍历key的切片，同时从中取出map的value
// 进阶玩法（比较难，如果没有写过，面试的时候就不要讲）：
// 1. 可以实现一个LinkedHashMap（双向链表+Map）结构，按照插入顺序维持住键值对的顺序，在增删场景更佳
// 2. 使用红黑树实现一个TreeMap结构，通过一个红黑树来维持键的顺序，适用需要按键排序的场景

// OrderedMap 写一个顶层接口，维持住所需要的方法，这里使用泛型实现
// K必须是comparable，用于可以比较，map的底层通常都是依赖hash表实现
// 哈希表的工作流程大致如下：
// 哈希 (Hashing)：当你向 map 中存入一个键值对（例如 myMap["name"] = "Alice"）时，
// Go 的运行时系统会先对键 "name" 计算一个哈希值（一个整数）。这个哈希值决定了这个键值对将要存放在底层数组的哪个位置（哪个“桶”里）。
// 比较 (Comparison)：当你用一个键去查找值时（例如 value := myMap["name"]），系统会：
// a. 再次对你提供的键 "name" 计算哈希值，找到对应的“桶”。
// b. 一个“桶”里可能因为哈希冲突存放了多个键值对。因此，系统需要遍历这个桶里的所有键，
// 用 == 运算符把你提供的键和桶里存储的键逐一比较，直到找到完全相等的那个键，然后返回它对应的值。
type OrderedMap[K comparable, V any] interface {
	// Set 设置一个键值对。
	// 如果键已存在，则更新其值，但顺序不变。
	// 如果键不存在，则添加到末尾。
	Set(key K, value V)

	// Get 根据键获取值。
	// 返回值和一个布尔值，布尔值表示键是否存在。
	Get(key K) (V, bool)
}

type OrderedMapBySlice[K comparable, V any] struct {
	lock sync.RWMutex
	data map[K]V
	keys []K
}
