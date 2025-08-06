package rng

import (
	"testing"
)

func TestBerlekampMassey(t *testing.T) {
	testfn := func(rng Interface, n int) []int {
		raw := make([]int, n)
		for i := 0; i < n; i++ {
			raw[i] = rng.Intn(2)
		}

		c := make([]int, n)
		b := make([]int, n)
		c[0], b[0] = 1, 1

		var L, m int
		var delta int
		var T []int

		for N := 0; N < n; N++ {

			delta = raw[N]
			for i := 1; i <= L; i++ {
				delta ^= c[i] & raw[N-i]
			}

			if delta == 1 {

				T = make([]int, n)
				copy(T, c)

				for i := 0; i < n-(N-m); i++ {
					c[N-m+i] ^= b[i]
				}

				if 2*L <= N {
					L = N + 1 - L
					m = N
					copy(b, T)
				}
			}
		}

		return c[:L+1]
	}
	testgroup := func(t *testing.T, rng Interface, n int) {
		t.Run("", func(t *testing.T) {
			lsfr := testfn(rng, n)
			// todo implement me
			_ = lsfr
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6)
	})
}
