package rng

import (
	"math/bits"
	"testing"
	"time"
)

func TestAvalanche(t *testing.T) {
	testfn := func(rng Interface, n int) float64 {
		type seederRNG interface {
			setSeed(seed int64)
			Uint64() uint64
		}
		var c int
		seed := time.Now().UnixNano()
		for i := 0; i < n; i++ {
			rng.Seed(seed)
			v0 := rng.Uint64()
			rng.Seed(seed ^ 1)
			v1 := rng.Uint64()

			c += bits.OnesCount64(v0 ^ v1)
		}
		return float64(c) / float64(n*64)
	}
	testgroup := func(t *testing.T, rng Interface, n int) {
		t.Run("", func(t *testing.T) {
			r := testfn(rng, n)
			if 0.5-r > 0.001 {
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
