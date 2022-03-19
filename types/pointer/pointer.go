package pointer

import (
	"sync/atomic"
	"unsafe"
)

type Pointer[T any] struct {
	value unsafe.Pointer
}

func New[T any]() Pointer[T] {
	return Pointer[T]{
		value: unsafe.Pointer(nil),
	}
}

func (p *Pointer[T]) Clone() Pointer[T] {
	np := New[T]()
	value := *p.Load()
	np.Store(&value)
	return np
}

func (p *Pointer[T]) Load() *T {
	pointer := atomic.LoadPointer(&p.value)
	return (*T)(pointer)
}

func (p *Pointer[T]) Store(pointer *T) {
	atomic.StorePointer(&p.value, unsafe.Pointer(pointer))
}

func (p *Pointer[T]) CompareAndSwap(old *T, new *T) bool {
	return atomic.CompareAndSwapPointer(&p.value, unsafe.Pointer(old), unsafe.Pointer(new))
}

func (p *Pointer[T]) Swap(new *T) *T {
	old := atomic.SwapPointer(&p.value, unsafe.Pointer(new))
	return (*T)(old)
}

func (p *Pointer[T]) IsNil() bool {
	return p.Load() == nil
}

func (p *Pointer[T]) Run(f func(*T)) bool {
	pointer := p.Load()
	if pointer == nil {
		return false
	}
	f(pointer)
	return true
}
