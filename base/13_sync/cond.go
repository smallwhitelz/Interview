package _3_sync

import "sync"

// Queue 可以改为泛型
type Queue struct {
	lock     *sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
	capacity int
	data     []any
}

func NewQueue(capacity int) *Queue {
	lock := &sync.Mutex{}
	return &Queue{
		lock:     lock,
		data:     make([]any, 0, capacity),
		notEmpty: sync.NewCond(lock),
		notFull:  sync.NewCond(lock),
		capacity: capacity,
	}
}

// Enqueue 有一个基本逻辑，就是如果已经满了，就要阻塞
func (q *Queue) Enqueue(data any) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.capacity == len(q.data) {
		// 如果已经满了，如何阻塞？

		// 在这里等待一个不满的信号
		q.notFull.Wait()

		// 你刚唤醒，结果别人把位置抢了
	}
	// 当你到这里的时候，可以断定这个位置被你抢到了
	q.data = append(q.data, data)
	// 唤醒一个等待不为空的信号
	q.notEmpty.Signal()
}

// Dequeue 基本逻辑，如果为空，就阻塞
func (q *Queue) Dequeue() any {
	// 用for不用if，因为也有可能那边刚唤醒这里，结果这个位置被别人抢了
	// 每次只保证一个G抢到执行
	q.lock.Lock()
	defer q.lock.Unlock()
	for len(q.data) == 0 {
		// 没有元素
		// 需要等一个不为空的信号
		q.notEmpty.Wait()
	}
	val := q.data[0]
	q.data = q.data[1:]
	// 你已经取走了一个元素，告诉对面的被阻塞的（可能有可能没有）
	q.notFull.Signal()
	return val
}
