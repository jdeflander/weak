// Package weak provides primitives for creating weak references to arbitrary values.
package weak

import (
	"runtime"
	"sync"
)

// Reference represents a weak reference to its corresponding value.
type Reference struct {
	mut *sync.RWMutex
	set *bool
	val *interface{}
}

// Get tries to retrieve the given reference's underlying value. It fails once the reference's corresponding value has
// been garbage collected.
func (r Reference) Get() (interface{}, bool) {
	r.mut.RLock()
	defer r.mut.RUnlock()
	return *r.val, *r.set
}

// Value represents a wrapper for an arbitrary value. It keeps track of garbage collection and updates its
// corresponding reference when necessary.
type Value struct {
	ref *Reference
}

// NewValue wraps the given value.
func NewValue(value interface{}) Value {
	set := true
	ref := &Reference{
		mut: new(sync.RWMutex),
		set: &set,
		val: &value,
	}

	runtime.SetFinalizer(ref, finalizer)
	return Value{ref: ref}
}

// Reference retrieves the given value's corresponding reference.
func (v Value) Reference() Reference {
	return *v.ref
}

func finalizer(ref *Reference) {
	ref.mut.Lock()
	defer ref.mut.Unlock()
	*ref.set = false
	*ref.val = nil
}
