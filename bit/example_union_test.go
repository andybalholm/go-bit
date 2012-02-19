package bit_test

import (
	bit "code.google.com/p/go-bit/bit"
	"fmt"
)

// Variadic Union function efficiently implemented with SetOr.
func Union(A ...*bit.Set) *bit.Set {
	S := new(bit.Set)
	for _, X := range A {
		S.SetOr(S, X)
	}
	return S
}

func ExampleSet_union() {
	A, B, C := bit.New(1, 2), bit.New(2, 3), bit.New(5)
	fmt.Println(Union(A, B, C))
	// Output: {1..3, 5}
}
