package main

import "fmt"

// 写一个大顶堆
// 完全二叉树（Complete Binary Tree）：除了最后一层外，每一层节点都被填满；并且最后一层的节点都集中在左边（没有“间隔”）。
// 重要性质：
// 它的节点按照“层次遍历”（level-order，或按层从左到右）填满。
// 因为没有“空洞”，可以用数组紧凑地按层顺序存放，每个节点的位置（索引）能映射到它的父/子节点的位置（索引）——这就是堆（binary heap）用数组实现的基础。
type Heap struct {
	items []int
}

// 0_base索引
//数组 [10, 20, 30, 40, 50, 60, 70]（索引 0..6）
//i=0（10）：left=1（20），right=2（30）
//i=1（20）：left=3（40），right=4（50）
//i=2（30）：left=5（60），right=6（70）

func NewHeap() *Heap {
	return &Heap{items: []int{}}
}

// parent 根据索引计算父节点的索引
func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) leftChild(index int) int {
	return index*2 + 1
}

func (h *Heap) rightChild(index int) int {
	return index*2 + 2
}

func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// 插入一个值
func (h *Heap) Insert(val int) {
	h.items = append(h.items, val)
	h.heapifyUp(len(h.items) - 1)
}

// 维护堆的属性：判断新加入的节点是否比父节点大，如果大的话就交换
func (h *Heap) heapifyUp(i int) {
	for i > 0 && h.items[i] > h.items[h.parent(i)] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

// 取出最大值
func (h *Heap) ExtractMax() (int, error) {
	if len(h.items) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	maxItem := h.items[0]
	lastIdx := len(h.items) - 1
	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]
	if len(h.items) > 0 {
		h.heapifyDown(0)
	}
	return maxItem, nil
}

// 下沉，取出最大值后，将最后一个元素放到顶部，然后和子节点判断，比他大就就让这个元素下沉
func (h *Heap) heapifyDown(i int) {
	lastIdx := len(h.items) - 1
	for {
		largest := i
		leftIdx := h.leftChild(i)
		rightIdx := h.rightChild(i)

		// 假如最大值在左边
		for leftIdx <= lastIdx && h.items[leftIdx] > h.items[largest] {
			largest = leftIdx
		}

		// 假如最大值在右边
		for rightIdx <= lastIdx && h.items[rightIdx] > h.items[largest] {
			largest = rightIdx
		}
		// 如果最大的索引还是在0处
		if largest == i {
			break
		}
		h.swap(i, largest)
		i = largest
	}
}

func (h *Heap) Peek() (int, error) {
	if len(h.items) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.items[0], nil
}

// Size returns the number of elements in the heap
func (h *Heap) Size() int {
	return len(h.items)
}

// IsEmpty checks if the heap is empty
func (h *Heap) IsEmpty() bool {
	return len(h.items) == 0
}

func (h *Heap) BuildHeap(arr []int) {
	h.items = make([]int, len(arr))
	copy(h.items, arr)
	for i := len(h.items)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *Heap) Print() {
	fmt.Println(h.items)
}

func main() {
	heap := NewHeap()
	arr := []int{5, 3, 8, 1, 9, 2}
	heap.BuildHeap(arr)
	fmt.Println(heap)
}
