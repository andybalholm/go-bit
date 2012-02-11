// Copyright 2012 Stefan Nilsson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bit_test

import (
	bit "code.google.com/p/go-bit/bit"
	"fmt"
)

// A = {0..99}; B = {0, 50..149, 200}
func ExampleSet() {
	A := new(bit.Set).AddRange(0, 100)
	B := bit.New(0, 200).AddRange(50, 150)
	fmt.Printf("A = %v; B = %v\n", A, B)
}

// sum({1..4}) = 10
func ExampleSet_Do() {
	A := bit.New(1, 2, 3, 4)
	sum := 0
	A.Do(func(n int) {
		sum += n
	})
	fmt.Printf("sum(%v) = %d\n", A, sum)
}
