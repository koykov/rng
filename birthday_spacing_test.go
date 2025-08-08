package rng

import (
	"math"
	"sort"
	"testing"
)

func TestBirthdaySpacing(t *testing.T) {
	testfn := func(rng Interface, n, d, samples int) float64 {
		lambda := math.Pow(float64(n), 3) / (4 * float64(d))
		k := 0

		for s := 0; s < samples; s++ {
			nums := make([]int, n)
			for i := range nums {
				nums[i] = rng.Intn(d)
			}

			sort.Ints(nums)
			spacings := make([]int, n-1)
			for i := 0; i < n-1; i++ {
				spacings[i] = nums[i+1] - nums[i]
			}

			for _, sp := range spacings {
				if sp == 0 {
					k++
				}
			}
		}

		observed := float64(k) / float64(samples)
		pValue := math.Exp(-lambda) * (1 + lambda)

		return math.Abs(pValue - observed)
	}
	testgroup := func(t *testing.T, rng Interface, n, d, samples int) {
		t.Run("", func(t *testing.T) {
			p := testfn(rng, n, d, samples)
			if p < 0.1 || p > 0.9 {
				t.Fail()
			}
		})
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1000, 1000000, 100)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1000, 1000000, 100)
	})
}
