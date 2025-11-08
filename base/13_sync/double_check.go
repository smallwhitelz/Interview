package _3_sync

import "sync"

type SafeMap struct {
	data map[string]string
	lock sync.RWMutex
}

// LoadOrStore 如果key存在，就返回老的数据，并且loaded为true
// 如果key不存在，就把key的值设为newVal，返回newVal，并且loaded为false
func (m *SafeMap) LoadOrStore(key string, newVal string) (val string, loaded bool) {
	// 粗暴玩法
	//m.lock.Lock()
	//defer m.lock.Unlock()

	// double check写法
	m.lock.RLock() // 读锁是共享锁，可以加多个
	oldVal, ok := m.data[key]
	if ok { // 读多写少，也就是大部分请求命中这个分支，就用double check写法
		m.lock.RUnlock()
		return oldVal, true
	}
	m.lock.RUnlock()
	m.lock.Lock()
	defer m.lock.Unlock()

	// 这个就是关键，再检查一次
	// double check 就是检查两次
	oldVal, ok = m.data[key]
	if ok {
		return oldVal, true
	}
	// 如果大部分请求直接是到这里，那就用直接写锁的玩法
	m.data[key] = newVal
	return newVal, false
	// 总结，所有的检查-做某事 类的问题，就是两种方案
	// 1. 直接写锁，检查，做某事。适合写多读少
	// 2. 先读锁，检查，符合条件返回；不符合条件，加写锁，再检查一遍，做某事。适合读多写少

	// 分布式环境下的检查-做某事也是一样的
}
