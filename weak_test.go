package weak_test

import (
	"fmt"
	"runtime"

	"github.com/jdeflander/weak"
)

func Example() {
	v := weak.NewValue(42)
	r := v.Reference()

	printReference(r)
	runtime.KeepAlive(v)

	runtime.GC()
	printReference(r)
	// Output:
	// 42 true
	// <nil> false
}

func printReference(reference weak.Reference) {
	i, ok := reference.Get()
	fmt.Println(i, ok)
}
