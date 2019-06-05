# Weak

Weak is a tiny Go package providing primitives for creating weak references.

## Installation

```
go get github.com/jdeflander/weak
```

## Usage

```go
package main

import (
	"fmt"
	"runtime"

	"github.com/jdeflander/weak"
)

func main() {
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
```
