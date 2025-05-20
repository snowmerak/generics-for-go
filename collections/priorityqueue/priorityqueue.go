package priorityqueue

import (
	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
)

// pqItem is an internal struct to hold an item and its priority.
type pqItem[T any, P comparable.Comparable[P]] struct {
	item     T
	priority P
}

// PriorityQueue implements a generic min-heap.
// T is the type of the item, and P is the type of the priority.
// P must satisfy comparable.Comparable.
type PriorityQueue[T any, P comparable.Comparable[P]] struct {
	items []pqItem[T, P]
}

// Interface defines the operations for a priority queue.
type Interface[T any, P comparable.Comparable[P]] interface {
	Push(item T, priority P)
	Pop() (T, P, bool)
	Peek() (T, P, bool)
	Len() int
	IsEmpty() bool
}

// New creates a new PriorityQueue.
func New[T any, P comparable.Comparable[P]]() *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		items: make([]pqItem[T, P], 0),
	}
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue[T, P]) Push(item T, priority P) {
	pq.items = append(pq.items, pqItem[T, P]{item: item, priority: priority})
	pq.heapifyUp(len(pq.items) - 1)
}

// heapifyUp maintains the min-heap property after adding an element.
func (pq *PriorityQueue[T, P]) heapifyUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		// If child's priority is less than parent's priority, swap them
		if pq.items[idx].priority.CompareTo(pq.items[parentIdx].priority) < 0 {
			pq.items[idx], pq.items[parentIdx] = pq.items[parentIdx], pq.items[idx]
			idx = parentIdx
		} else {
			break // Heap property is satisfied
		}
	}
}

// Pop removes and returns the item with the highest priority (lowest priority value).
// Returns the item, its priority, and true if successful.
// Returns zero values for item and priority, and false if the queue is empty.
func (pq *PriorityQueue[T, P]) Pop() (T, P, bool) {
	if pq.IsEmpty() {
		var zeroT T
		var zeroP P
		return zeroT, zeroP, false
	}

	n := len(pq.items)
	itemToPop := pq.items[0] // Item with highest priority (min value) is at the root

	// Move the last element to the root
	pq.items[0] = pq.items[n-1]
	pq.items = pq.items[:n-1] // Shorten the slice

	if len(pq.items) > 0 {
		pq.heapifyDown(0) // Restore heap property
	}

	return itemToPop.item, itemToPop.priority, true
}

// heapifyDown maintains the min-heap property after removing the root element.
func (pq *PriorityQueue[T, P]) heapifyDown(idx int) {
	n := len(pq.items)
	for {
		smallest := idx
		leftIdx := 2*idx + 1
		rightIdx := 2*idx + 2

		// Check if left child exists and is smaller than current smallest
		if leftIdx < n && pq.items[leftIdx].priority.CompareTo(pq.items[smallest].priority) < 0 {
			smallest = leftIdx
		}

		// Check if right child exists and is smaller than current smallest
		if rightIdx < n && pq.items[rightIdx].priority.CompareTo(pq.items[smallest].priority) < 0 {
			smallest = rightIdx
		}

		// If the smallest is not the current element, swap them and continue
		if smallest != idx {
			pq.items[idx], pq.items[smallest] = pq.items[smallest], pq.items[idx]
			idx = smallest
		} else {
			break // Heap property is satisfied
		}
	}
}

// Peek returns the item with the highest priority (lowest priority value) without removing it.
// Returns the item, its priority, and true if successful.
// Returns zero values for item and priority, and false if the queue is empty.
func (pq *PriorityQueue[T, P]) Peek() (T, P, bool) {
	if pq.IsEmpty() {
		var zeroT T
		var zeroP P
		return zeroT, zeroP, false
	}
	// In a min-heap, the item with the highest priority (lowest value) is at the root.
	// This is correct as Push and Pop maintain the heap property.
	return pq.items[0].item, pq.items[0].priority, true
}

// Len returns the number of items in the priority queue.
func (pq *PriorityQueue[T, P]) Len() int {
	return len(pq.items)
}

// IsEmpty checks if the priority queue is empty.
func (pq *PriorityQueue[T, P]) IsEmpty() bool {
	return len(pq.items) == 0
}
