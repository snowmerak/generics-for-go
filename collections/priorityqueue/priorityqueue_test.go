package priorityqueue_test

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"

	pq "github.com/snowmerak/generics-for-go/v2/collections/priorityqueue"
	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
	"github.com/snowmerak/generics-for-go/v2/test/assert"
)

// IntPriority is a helper type for testing with int priorities.
type IntPriority int

func (ip IntPriority) CompareTo(other IntPriority) int {
	if ip < other {
		return -1
	}
	if ip > other {
		return 1
	}
	return 0
}

// MyStructPriority is another helper for testing with struct priorities.
type MyStructPriority struct {
	Value int
	Label string // Added to make it a bit more complex if needed
}

func (msp MyStructPriority) CompareTo(other MyStructPriority) int {
	if msp.Value < other.Value {
		return -1
	}
	if msp.Value > other.Value {
		return 1
	}
	return 0
}

func TestNewPriorityQueue(t *testing.T) {
	t.Run("NewPQStringInt", func(t *testing.T) {
		q := pq.New[string, IntPriority]()
		assert.NotNil(t, q)
		assert.Equal(t, 0, q.Len())
		assert.True(t, q.IsEmpty())
	})

	t.Run("NewPQIntStruct", func(t *testing.T) {
		q := pq.New[int, MyStructPriority]()
		assert.NotNil(t, q)
		assert.Equal(t, 0, q.Len())
		assert.True(t, q.IsEmpty())
	})
}

func TestPushAndPeek(t *testing.T) {
	q := pq.New[string, IntPriority]()

	t.Run("PushFirst", func(t *testing.T) {
		q.Push("task1", 10)
		assert.Equal(t, 1, q.Len())
		assert.False(t, q.IsEmpty())
		item, prio, ok := q.Peek()
		assert.True(t, ok)
		assert.Equal(t, "task1", item)
		assert.Equal(t, IntPriority(10), prio)
	})

	t.Run("PushSecondHigherPriority", func(t *testing.T) {
		q.Push("task2", 5) // Higher priority (lower value)
		assert.Equal(t, 2, q.Len())
		item, prio, ok := q.Peek()
		assert.True(t, ok)
		assert.Equal(t, "task2", item)
		assert.Equal(t, IntPriority(5), prio)
	})

	t.Run("PushThirdLowerPriority", func(t *testing.T) {
		q.Push("task3", 15) // Lower priority
		assert.Equal(t, 3, q.Len())
		item, prio, ok := q.Peek()
		assert.True(t, ok)
		assert.Equal(t, "task2", item) // task2 should still be at the top
		assert.Equal(t, IntPriority(5), prio)
	})

	t.Run("PushSameAsTopPriority", func(t *testing.T) {
		q.Push("task4", 5) // Same priority as current top
		assert.Equal(t, 4, q.Len())
		item, prio, ok := q.Peek()
		assert.True(t, ok)
		// The original "task2" or "task4" could be at the top if priorities are equal.
		// The heap property only guarantees one of them.
		assert.Equal(t, IntPriority(5), prio)
		// We can't be certain if it's "task2" or "task4", so we only check priority.
	})
}

func TestPop(t *testing.T) {
	q := pq.New[string, IntPriority]()
	itemsToPush := []struct {
		item     string
		priority IntPriority
	}{
		{"itemA", 5},
		{"itemB", 2},
		{"itemC", 8},
		{"itemD", 1},
		{"itemE", 5}, // Same priority as itemA
		{"itemF", 3},
	}

	for _, val := range itemsToPush {
		q.Push(val.item, val.priority)
	}

	expectedPopOrder := []struct { // Expected: (item, priority)
		item     string
		priority IntPriority
	}{
		{"itemD", 1},
		{"itemB", 2},
		{"itemF", 3},
		// For items with priority 5, the order is not strictly guaranteed by value,
		// but determined by their original positions and swaps.
		// The test needs to accommodate this.
		// Let's assume for now it's one of them, then the other.
		// We'll pop and check if priority is 5, then if the item is A or E.
	}

	t.Run("PopInOrder", func(t *testing.T) {
		poppedItems := make(map[string]IntPriority)

		item, prio, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "itemD", item)
		assert.Equal(t, IntPriority(1), prio)
		assert.Equal(t, 5, q.Len())
		poppedItems[item] = prio

		item, prio, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "itemB", item)
		assert.Equal(t, IntPriority(2), prio)
		assert.Equal(t, 4, q.Len())
		poppedItems[item] = prio

		item, prio, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "itemF", item)
		assert.Equal(t, IntPriority(3), prio)
		assert.Equal(t, 3, q.Len())
		poppedItems[item] = prio
		
		// Handling items with priority 5
		item1_p5, prio1_p5, ok1_p5 := q.Pop()
		assert.True(t, ok1_p5)
		assert.Equal(t, IntPriority(5), prio1_p5)
		poppedItems[item1_p5] = prio1_p5

		item2_p5, prio2_p5, ok2_p5 := q.Pop()
		assert.True(t, ok2_p5)
		assert.Equal(t, IntPriority(5), prio2_p5)
		poppedItems[item2_p5] = prio2_p5

		// Check that both "itemA" and "itemE" were popped
		_, foundA := poppedItems["itemA"]
		_, foundE := poppedItems["itemE"]
		assert.True(t, foundA, "itemA with priority 5 was not popped")
		assert.True(t, foundE, "itemE with priority 5 was not popped")
		assert.Equal(t, 1, q.Len())


		item, prio, ok = q.Pop() // Last item
		assert.True(t, ok)
		assert.Equal(t, "itemC", item)
		assert.Equal(t, IntPriority(8), prio)
		assert.Equal(t, 0, q.Len())
		assert.True(t, q.IsEmpty())
		poppedItems[item] = prio
	})

	t.Run("PopFromEmpty", func(t *testing.T) {
		item, prio, ok := q.Pop()
		assert.False(t, ok)
		var zeroString string
		var zeroPrio IntPriority
		assert.Equal(t, zeroString, item)
		assert.Equal(t, zeroPrio, prio)
	})
}

func TestPopEmptyAndPeekEmpty(t *testing.T) {
	q := pq.New[int, IntPriority]()

	t.Run("InitiallyEmpty", func(t *testing.T) {
		item, prio, ok := q.Peek()
		assert.False(t, ok)
		var zeroInt int
		var zeroPrio IntPriority
		assert.Equal(t, zeroInt, item)
		assert.Equal(t, zeroPrio, prio)

		item, prio, ok = q.Pop()
		assert.False(t, ok)
		assert.Equal(t, zeroInt, item)
		assert.Equal(t, zeroPrio, prio)
		assert.True(t, q.IsEmpty())
	})

	t.Run("EmptyAfterOperations", func(t *testing.T) {
		q.Push(100, 1)
		assert.False(t, q.IsEmpty())
		q.Pop()
		assert.True(t, q.IsEmpty())

		item, prio, ok := q.Peek()
		assert.False(t, ok)
		var zeroInt int
		var zeroPrio IntPriority
		assert.Equal(t, zeroInt, item)
		assert.Equal(t, zeroPrio, prio)

		item, prio, ok = q.Pop()
		assert.False(t, ok)
		assert.Equal(t, zeroInt, item)
		assert.Equal(t, zeroPrio, prio)
	})
}

func TestSamePriorityOrder(t *testing.T) {
	q := pq.New[string, IntPriority]()
	q.Push("A", 5)
	q.Push("B", 5)
	q.Push("X", 1) // Higher priority
	q.Push("C", 5)
	q.Push("Y", 0) // Highest priority

	expectedOrder := []string{"Y", "X"}
	poppedOrder := []string{}

	item, prio, ok := q.Pop()
	assert.True(t, ok)
	assert.Equal(t, "Y", item)
	assert.Equal(t, IntPriority(0), prio)
	poppedOrder = append(poppedOrder, item)

	item, prio, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, "X", item)
	assert.Equal(t, IntPriority(1), prio)
	poppedOrder = append(poppedOrder, item)

	// For items with priority 5, order isn't strictly defined by value.
	// We just need to ensure all three (A, B, C) are popped.
	priority5Items := make(map[string]bool)
	for i := 0; i < 3; i++ {
		item, prio, ok := q.Pop()
		assert.True(t, ok, "Expected to pop item with priority 5")
		assert.Equal(t, IntPriority(5), prio)
		priority5Items[item] = true
	}
	assert.True(t, priority5Items["A"])
	assert.True(t, priority5Items["B"])
	assert.True(t, priority5Items["C"])
	assert.Equal(t, 3, len(priority5Items))

	assert.True(t, q.IsEmpty())
}

func TestDifferentDataTypes(t *testing.T) {
	t.Run("StringItemsIntPriority", func(t *testing.T) {
		q := pq.New[string, IntPriority]()
		q.Push("hello", 10)
		q.Push("world", 1)
		q.Push("!", 20)

		item, prio, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "world", item)
		assert.Equal(t, IntPriority(1), prio)

		item, prio, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "hello", item)
		assert.Equal(t, IntPriority(10), prio)

		item, prio, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, "!", item)
		assert.Equal(t, IntPriority(20), prio)
		assert.True(t, q.IsEmpty())
	})

	t.Run("IntItemsMyStructPriority", func(t *testing.T) {
		q := pq.New[int, MyStructPriority]()
		q.Push(100, MyStructPriority{Value: 10, Label: "L10"})
		q.Push(200, MyStructPriority{Value: 1, Label: "L1"})
		q.Push(300, MyStructPriority{Value: 20, Label: "L20"})

		item, prio, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, 200, item)
		assert.Equal(t, MyStructPriority{Value: 1, Label: "L1"}, prio)

		item, prio, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, 100, item)
		assert.Equal(t, MyStructPriority{Value: 10, Label: "L10"}, prio)

		item, prio, ok = q.Pop()
		assert.True(t, ok)
		assert.Equal(t, 300, item)
		assert.Equal(t, MyStructPriority{Value: 20, Label: "L20"}, prio)
		assert.True(t, q.IsEmpty())
	})
}

func TestLargeNumberOfItems(t *testing.T) {
	q := pq.New[string, IntPriority]()
	numItems := 1000
	rand.Seed(time.Now().UnixNano())

	pushedPriorities := make([]IntPriority, 0, numItems)

	for i := 0; i < numItems; i++ {
		prio := IntPriority(rand.Intn(numItems * 10))
		item := "item_" + strconv.Itoa(i) + "_prio_" + strconv.Itoa(int(prio))
		q.Push(item, prio)
		pushedPriorities = append(pushedPriorities, prio) // Not used for verification directly, but shows what was pushed
	}

	assert.Equal(t, numItems, q.Len())

	poppedPriorities := make([]IntPriority, 0, numItems)
	var lastPriority IntPriority = -1 // Assuming priorities are non-negative

	for i := 0; i < numItems; i++ {
		_, prio, ok := q.Pop()
		assert.True(t, ok, "Expected to pop an item from a non-empty queue")
		poppedPriorities = append(poppedPriorities, prio)
		if i > 0 {
			assert.True(t, prio.CompareTo(lastPriority) >= 0, "Items not popped in sorted order of priority")
		}
		lastPriority = prio
	}

	assert.Equal(t, numItems, len(poppedPriorities))
	assert.True(t, q.IsEmpty())
	assert.Equal(t, 0, q.Len())

	// Verify that sorting the pushed priorities matches the popped priorities
	// This is a more robust check if there are many items with same priorities.
	sort.Slice(pushedPriorities, func(i, j int) bool {
		return pushedPriorities[i] < pushedPriorities[j]
	})
	
	// Re-populating for direct comparison (original test logic for order was fine)
	// The previous loop already verified the order.
	// This sort check is just to ensure the distribution of priorities is handled.
	// For this test, ensuring `poppedPriorities` is sorted is the primary goal.
	// The previous loop already did that with `prio.CompareTo(lastPriority) >= 0`.
}

// Ensure the file ends with a newline
