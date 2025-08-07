package rng

import (
	"strconv"
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
		var delta, mask int

		for nIter := 0; nIter < n; nIter++ {
			delta = raw[nIter]
			for i := 1; i <= L; i++ {
				delta ^= c[i] & raw[nIter-i]
			}

			if delta == 1 {
				t_ := make([]int, n)
				copy(t_, c)

				mask = 1
				shift := nIter - m
				for i := 0; i < n; i++ {
					if i+shift < n {
						c[i+shift] ^= b[i] & mask
					}
				}

				if L <= nIter/2 {
					L = nIter + 1 - L
					m = nIter
					copy(b, t_)
				}
			}
		}

		res := make([]int, L+1)
		copy(res, c[:L+1])
		return res
	}
	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				lsfr := testfn(rng, step)
				rate := float64(len(lsfr)) / float64(step)
				if rate <= 0.5 {
					t.Errorf("LSFR rate %f too small", rate)
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 10, 100, 1000, 10000)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 10, 100, 1000, 10000)
	})
}
