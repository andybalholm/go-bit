package bit_test

import (
	. "bit"
	"testing"
)

func BenchmarkMinPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinPos(0xcafecafecafecafe)
	}
}

func BenchmarkMaxPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxPos(0xcafecafecafecafe)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(0xcafecafecafecafe)
	}
}

// Number of words in test set.
const nw = 1 << 10

func BenchmarkSize(b *testing.B) {
	b.StopTimer()
	S := BuildTestSet(nw << 3) // Allocates nw<<3 bytes = nw words.
	b.StartTimer()

	for i := 0; i < b.N/nw; i++ { // Measure time per word.
		S.Size()
	}
}

func BenchmarkSetAnd(b *testing.B) {
	b.StopTimer()
	S := New().SetWord(nw-1, 1).Clear() // Allocates nw words.
	A := BuildTestSet(nw << 3)
	B := BuildTestSet(nw << 3)
	b.StartTimer()

	for i := 0; i < b.N/nw; i++ { // Measure time per word.
		S.SetAnd(A, B)
	}
}

func BenchmarkDo(b *testing.B) {
	b.StopTimer()
	S := BuildTestSet(b.N) // As Do is pretty fast, S can be pretty big.
	b.StartTimer()

	S.Do(func(i int) {})
}

func BenchmarkSetWord(b *testing.B) {
	b.StopTimer()
	S := New().SetWord(MaxInt>>6, 1)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		S.SetWord(i&(MaxInt>>6), 1)
	}
}

func BenchmarkRemoveMin(b *testing.B) {
	b.StopTimer()
	S := BuildTestSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		S.RemoveMin()
	}
}

func BenchmarkRemoveMax(b *testing.B) {
	b.StopTimer()
	S := BuildTestSet(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		S.RemoveMax()
	}
}

// Quickly builds a set of n somewhat random elements from 0..8n-1.
func BuildTestSet(n int) *Set {
	S := New().Add(8*n - 1).Clear() // Allocates n bytes.

	lfsr := uint16(0xace1) // linear feedback shift register
	for i := 0; i < n; i++ {
		bit := (lfsr>>0 ^ lfsr>>2 ^ lfsr>>3 ^ lfsr>>5) & 1
		lfsr = lfsr>>1 | bit<<15
		e := i<<3 + int(lfsr&0x7)
		S.Add(e) // Add a number from 8i..8i+7.
	}
	return S
}
