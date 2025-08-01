package rng

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestKolmogorovSmirnov(t *testing.T) {
	testfn := func(rng Interface, n int) (r float64) {
		raw := make([]float64, n)
		for i := 0; i < n; i++ {
			raw[i] = rng.Float64()
		}

		sorted := make([]float64, n)
		copy(sorted, raw)
		sort.Float64s(sorted)

		for i := 0; i < n; i++ {
			empirical := float64(i+1) / float64(n)
			theoretical := sorted[i]
			diff := math.Abs(empirical - theoretical)
			if diff > r {
				r = diff
			}
		}
		return
	}
	testgroup := func(t *testing.T, rng Interface, n int, alphas ...float64) {
		for _, alpha := range alphas {
			t.Run(fmt.Sprintf("%f", alpha), func(t *testing.T) {
				x := -math.Log(alpha/2) / 2
				Dcrit := math.Sqrt(x)
				Dn := testfn(rng, n)
				if Dn >= Dcrit {
					t.Errorf("%f should be less than %f", Dn, Dcrit)
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6, 0.05, 0.01)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6, 0.05, 0.01)
	})
}
