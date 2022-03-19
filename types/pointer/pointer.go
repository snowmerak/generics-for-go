// pointer is a pointer to a value of type T.
package pointer

import (
	"sync/atomic"
	"unsafe"
)

// Pointer is a container of T pointer.
type Pointer[T any] struct {
	value unsafe.Pointer
}

// New returns a new Pointer.
func New[T any](pointer *T) Pointer[T] {
	return Pointer[T]{
		value: unsafe.Pointer(pointer),
	}
}

// Clone returns a new Pointer with the same value as the old one.
func (p *Pointer[T]) Clone() Pointer[T] {
	value := *p.Load()
	return New(&value)
}

// Load returns the value of the Pointer.
func (p *Pointer[T]) Load() *T {
	pointer := atomic.LoadPointer(&p.value)
	return (*T)(pointer)
}

// Store sets the value of the Pointer.
func (p *Pointer[T]) Store(pointer *T) {
	atomic.StorePointer(&p.value, unsafe.Pointer(pointer))
}

// CompareAndSwap sets the value of the Pointer if the current value is equal to the expected value.
func (p *Pointer[T]) CompareAndSwap(old *T, new *T) bool {
	return atomic.CompareAndSwapPointer(&p.value, unsafe.Pointer(old), unsafe.Pointer(new))
}

// Swap sets the value of the Pointer and returns the previous value.
func (p *Pointer[T]) Swap(new *T) *T {
	old := atomic.SwapPointer(&p.value, unsafe.Pointer(new))
	return (*T)(old)
}

// IsNil returns true if the Pointer is nil.
func (p *Pointer[T]) IsNil() bool {
	return p.Load() == nil
}

// Run is a function that takes a pointer of type T.
// if pointer is nil, returns false.
func (p *Pointer[T]) Run(f func(*T)) bool {
	pointer := p.Load()
	if pointer == nil {
		return false
	}
	f(pointer)
	return true
}
