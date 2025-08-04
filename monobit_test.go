package rng

import (
	"math"
	"math/bits"
	"testing"
)

func TestMonobit(t *testing.T) {
	testfn := func(rng Interface, n int) float64 {
		var n0, n1 float64
		for i := 0; i < n; i++ {
			v := rng.Uint64()
			o := bits.OnesCount64(v)
			n1 += float64(o)
			n0 += float64(64 - o)
		}
		return math.Min(n0, n1) / math.Max(n0, n1)
	}
	testgroup := func(t *testing.T, rng Interface, n int) {
		t.Run("", func(t *testing.T) {
			rate := testfn(rng, n)
			if 1.0-rate > 0.01 {
				t.Fail()
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
