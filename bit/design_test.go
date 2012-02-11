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
	"math/rand"
	"testing"
)

const RunExamples = false

// A few examples demonstrating the bits package API
func TestDesign(t *testing.T) {
	if !RunExamples {
		return
	}

	// Building and printing sets
	A := new(bit.Set).AddRange(0, 100)     // A = {0..99}
	B := bit.New(0, 200).AddRange(50, 150) // B = {0, 50..149, 200}
	fmt.Printf("A = %v; B = %v\n", A, B)

	// Set operators
	S := A.Xor(B)                    // S = A ∆ B
	C := A.Or(B).AndNot(A.And(B))    // C = (A ∪ B) ∖ (A ∩ B)
	D := A.AndNot(B).Or(B.AndNot(A)) // D = (A ∖ B) ∪ (B ∖ A)
	if C.Equals(S) && D.Equals(S) {
		fmt.Printf("A ∆ B = %v\n", S)
	}

	// Iteration
	sum := 0
	S.Do(func(n int) {
		sum += n
	})
	fmt.Printf("sum(A ∆ B) = %d\n", sum)

	// Variadic Union function efficiently implemented with SetOr
	Union := func(A ...*bit.Set) *bit.Set {
		S := new(bit.Set)
		for _, X := range A {
			S.SetOr(S, X)
		}
		return S
	}
	fmt.Printf("A ∪ B ∪ (A ∆ B) = %v\n", Union(A, B, S))

	// Creating a set of primes using Sieve of Eratosthenes
	Sieve, Primes := new(bit.Set).AddRange(2, 50), new(bit.Set)
	for !Sieve.IsEmpty() {
		p := Sieve.RemoveMin()
		Primes.Add(p)
		Sieve.Do(func(n int) {
			if n%p == 0 {
				Sieve.Remove(n)
			}
		})
	}
	fmt.Printf("\n%v\n", Primes)

	// Constructing sets from words
	const faraway = 46                         // billion light years
	Universe := bit.New().AddRange(0, faraway) // Universe = {0..faraway-1}
	Even := bit.New().SetWord(0, 1<<faraway/3) // Even = {0, 2, 4, ..., faraway-2}
	Odd := Universe.AndNot(Even)
	Even.FlipRange(0, faraway)
	if !Even.Equals(Odd) {
		t.Errorf("Even isn't odd; want odd.")
	}

	// Memory management
	S = bit.New(100, 100000) // Make a set that occupies a few kilobyes.
	S.RemoveMax()            // The large element is gone. Memory remains.
	S = bit.New().Set(S)     // Give excess capacity to garbage collector.

	// Memory hints and memory reuse
	IntegerRuns(10, 5)

	// Computing suitable allocation size (using MaxPos and MaxInt)
	// favoring powers of two and guaranteeing linear amortized cost
	// for a repeated number of allocations.
	newSize := func(want, had int) int {
		return max(want, NextPowerOfTwo(had))
	}

	say := func(h, w, g int) {
		fmt.Printf("Had %#x, wanted %#x, got %#x.\n", h, w, g)
	}

	had := 0xff0
	want := had + 0x008       // shy of next power of two
	got := newSize(want, had) // got == NextPowerOfTwo(had)
	say(had, want, got)

	had = got
	want = had + 0x2000      // overshooting next power of two
	got = newSize(want, had) // got == want
	say(had, want, got)

	had = got
	want = had + 0x1000      // hitting next power of two
	got = newSize(want, had) // got == want
	say(had, want, got)
}

// NextPowerOfTwo returns the smallest p = 1, 2, 4, ..., 2^k such that p > n,
// or MaxInt if p > MaxInt.
func NextPowerOfTwo(n int) (p int) {
	if n <= 0 {
		return 1
	}

	if k := bit.MaxPos(uint64(n)) + 1; k < bit.BitsPerWord-1 {
		return 1 << uint(k)
	}
	return bit.MaxInt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IntegerRuns(s, n) generates random sets R(i), i = 1, 2, ..., n,
// with elements drawn from 0..i*s-1 and computes the number of runs
// in each set. A run is a sequence of at least three integers
// a, a+1, ..., b, such that {a..b} ⊆ S and {a-1, b+1} ∩ S = ∅.
func IntegerRuns(start, n int) {
	// Give a capacity hint.
	R := new(bit.Set).Add(n*start - 1).Clear()

	fmt.Printf("\n%8s %8s %8s\n", "Max", "Size", "Runs")
	for i := 1; i <= n; i++ {
		// Reuse memory from last iteration.
		R.Clear()

		// Create a random set with ~86% population.
		max := i * start
		for i := 0; i < 2*max; i++ {
			R.Add(rand.Intn(max))
		}

		// Compute the number of runs.
		runs, length, prev := 0, 0, -2
		R.Do(func(i int) {
			if i == prev+1 { // Continue run.
				length++
			} else { // Start a new run.
				if length >= 3 {
					runs++
				}
				length = 1
			}
			prev = i
		})
		if length >= 3 {
			runs++
		}
		fmt.Printf("%8d %8d %8d", max, R.Size(), runs)
		if max <= 32 {
			fmt.Printf("%4s%v", "", R)
		}
		fmt.Println()
	}
	fmt.Println()
}
