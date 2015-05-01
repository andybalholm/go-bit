This package contains a comprehensive bit set implementation and some utility bit functions for uint64 words. The code is intended to be fast and memory efficient.

Even if bit fiddling doesn't rock your boat, there are some code snippets you might enjoy:

- The implementation-specific constants (bit.MaxInt, etc) are computed using the << operator. This is a trick to make them untyped.

- bit.Count is a fun algorithm for counting the number of bits in a word.

- bit.MinPos is another fun algorithm that uses a De Bruijn sequence to compute the minimum nonzero bit in a word.

- Panics() in funcs\_test.go is a general purpose function using reflection to test if functions panic as intended. It did help me find some insidious bugs.