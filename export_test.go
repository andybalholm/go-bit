package bit

func (S *Set) Sdata() []uint64             { return S.data }
func (S *Set) Smin() int                   { return S.min }
func FindMinFrom(n int, data []uint64) int { return findMinFrom(n, data) }
func NextPow2(n int) (p int)               { return nextPow2(n) }
