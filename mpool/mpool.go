package mpool

import (
	"runtime"
	"sync"
	"unsafe"
)

type MemoryPool[T any] struct {
	pool      []unsafe.Pointer
	slicePool map[int][]unsafe.Pointer
	lock      sync.Mutex
}

func New[T any]() *MemoryPool[T] {
	return &MemoryPool[T]{
		pool:      make([]unsafe.Pointer, 0, 10),
		slicePool: make(map[int][]unsafe.Pointer),
		lock:      sync.Mutex{},
	}
}

func (mp *MemoryPool[T]) Get() *T {
	mp.lock.Lock()
	defer mp.lock.Unlock()
	if len(mp.pool) == 0 {
		p := new(T)
		return p
	}
	p := mp.pool[len(mp.pool)-1]
	mp.pool = mp.pool[:len(mp.pool)-1]
	return (*T)(p)
}

func (mp *MemoryPool[T]) GetSlice(size int) []T {
	mp.lock.Lock()
	defer mp.lock.Unlock()
	if slices, ok := mp.slicePool[size]; ok {
		if len(slices) == 0 {
			r := make([]T, size)
			return r
		}
		r := slices[len(slices)-1]
		slices = slices[:len(slices)-1]
		mp.slicePool[size] = slices
		return *(*[]T)(r)
	}
	r := make([]T, size)
	return r
}

func (mp *MemoryPool[T]) Put(v *T) {
	mp.lock.Lock()
	defer mp.lock.Unlock()
	mp.pool = append(mp.pool, unsafe.Pointer(v))
}

func (mp *MemoryPool[T]) PutSlice(v *[]T) {
	mp.lock.Lock()
	defer mp.lock.Unlock()
	if slices, ok := mp.slicePool[len(*v)]; ok {
		slices = append(slices, unsafe.Pointer(v))
		mp.slicePool[len(*v)] = slices
		return
	}
	mp.slicePool[len(*v)] = []unsafe.Pointer{unsafe.Pointer(v)}
}

func (mp *MemoryPool[T]) Free(num int) {
	mp.lock.Lock()
	defer mp.lock.Unlock()
	if len(mp.pool) > num {
		mp.pool = mp.pool[:num]
	} else {
		mp.pool = nil
	}
	runtime.GC()
}

func (mp *MemoryPool[T]) FreeSlice(num int, size int) {
	mp.lock.Lock()
	defer mp.lock.Unlock()
	if slices, ok := mp.slicePool[size]; ok {
		if len(slices) > num {
			slices = slices[:num]
		} else {
			slices = nil
		}
		mp.slicePool[size] = slices
	}
	runtime.GC()
}

func (mp *MemoryPool[T]) FreeAll() {
	mp.lock.Lock()
	defer mp.lock.Unlock()
	mp.pool = nil
	mp.slicePool = nil
	runtime.GC()
}
