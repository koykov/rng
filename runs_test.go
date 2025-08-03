package rng

import (
	"math"
	"math/bits"
	"testing"
	"unsafe"
)

func TestRuns(t *testing.T) {
	testfn := func(rng Interface, n int) (Z, pValue float64) {
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

		Z = (R - mu) / sigma
		pValue = 2 * (1 - math.Erf(math.Abs(Z)/math.Sqrt2))
		return
	}
	testgroup := func(t *testing.T, rng Interface, n int) {
		t.Run("", func(t *testing.T) {
			const limit = 0.05
			_, pValue := testfn(rng, n)
			if pValue < limit {
				t.Errorf("pValue (%f) must be less than %f", pValue, limit)
			}
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6)
	})
}
