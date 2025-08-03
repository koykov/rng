package rng

import (
	"math"
	"math/bits"
	"testing"
	"unsafe"
)

func TestMonobit(t *testing.T) {
	testfn := func(rng Interface, n int) float64 {
		var n0, n1 float64
		for i := 0; i < n; i++ {
			v := rng.Float64()
			u := *(*uint64)(unsafe.Pointer(&v))
			t := bits.OnesCount64(u)
			n1 += float64(t)
			n0 += float64(64 - t)
		}
		return math.Min(n0, n1) / math.Max(n0, n1)
	}
	testgroup := func(t *testing.T, rng Interface, n int) {
		t.Run("", func(t *testing.T) {
			rate := testfn(rng, n)
			t.Log(rate) // todo check
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6)
	})
}
