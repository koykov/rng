package rng

import (
	"math"
	"math/bits"
	"testing"
	"unsafe"
)

func TestRuns(t *testing.T) {
	testfn := func(rng Interface, n int) (r float64) {
		var (
			n0, n1, R float64
			b, pb     uint64
		)
		for i := 0; i < n; i++ {
			v := rng.Float64()
			u := *(*uint64)(unsafe.Pointer(&v))
			n1 += float64(bits.OnesCount64(u))
			for j := 0; j < 64; j++ {
				b = (u >> j) & 1
				if b != pb {
					R++
				}
				pb = b
			}
		}
		n0 = float64(n) - n1
		if n0 == n1 {
			return // all bits are equal
		}

		mu := (2*n1*n0)/(n1+n0) + 1
		sigma := math.Sqrt((2 * n1 * n0 * (2*n1*n0 - n1 - n0)) / ((n1 + n0) * math.Sqrt(n1+n0-1)))

		Z := (R - mu) / sigma
		pValue := 2 * (1 - math.Erf(math.Abs(Z)/math.Sqrt2))

		return pValue
	}
	var _ = testfn
}
