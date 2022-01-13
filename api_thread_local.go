package routine

import "sync/atomic"

var (
	threadLocalIndex            int32 = -1
	inheritableThreadLocalIndex int32 = -1
)

// ThreadLocal provides goroutine-local variables.
type ThreadLocal interface {
	// Id returns the global id of instance
	Id() int

	// Get returns the value in the current goroutine's local threadLocalImpl, if it was set before.
	Get() Any

	// Set copy the value into the current goroutine's local threadLocalImpl, and return the old value.
	Set(value Any)

	// Remove delete the value from the current goroutine's local threadLocalImpl, and return it.
	Remove()
}

type Supplier = func() Any

// NewThreadLocal create and return a new ThreadLocal instance.
func NewThreadLocal() ThreadLocal {
	return &threadLocalImpl{id: int(atomic.AddInt32(&threadLocalIndex, 1))}
}

// NewThreadLocalWithInitial create and return a new ThreadLocal instance. The initial value is determined by invoking the supplier method.
func NewThreadLocalWithInitial(supplier Supplier) ThreadLocal {
	return &threadLocalImpl{id: int(atomic.AddInt32(&threadLocalIndex, 1)), supplier: supplier}
}

// NewInheritableThreadLocal create and return a new ThreadLocal instance.
func NewInheritableThreadLocal() ThreadLocal {
	return &inheritableThreadLocalImpl{id: int(atomic.AddInt32(&inheritableThreadLocalIndex, 1))}
}

// NewInheritableThreadLocalWithInitial create and return a new ThreadLocal instance. The initial value is determined by invoking the supplier method.
func NewInheritableThreadLocalWithInitial(supplier Supplier) ThreadLocal {
	return &inheritableThreadLocalImpl{id: int(atomic.AddInt32(&inheritableThreadLocalIndex, 1)), supplier: supplier}
}
