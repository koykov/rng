package rng

import (
	"math"
	"strconv"
	"testing"
)

func TestAutocorrelation(t *testing.T) {
	testfn := func(rng Interface, n, lag int) float64 {
		if lag >= n {
			return 0
		}
		raw := make([]float64, n)
		var mean float64
		for i := 0; i < n; i++ {
			raw[i] = rng.Float64()
			mean += raw[i]
		}
		mean /= float64(n)

		var numerator, denominator float64
		for i := 0; i < n-lag; i++ {
			numerator += (raw[i] - mean) * (raw[i+lag] - mean)
			denominator += math.Pow(raw[i]-mean, 2)
		}

		return numerator / denominator
	}

	testgroup := func(t *testing.T, rng Interface, n int, lags ...int) {
		for _, lag := range lags {
			t.Run(strconv.Itoa(lag), func(t *testing.T) {
				limit := 1.96 / math.Sqrt(float64(n))
				v := testfn(rng, n, lag)
				if v <= -limit || v >= limit {
					t.Errorf("value %f out of range [-%f, %f]", v, limit, limit)
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6, 1, 10, 100, 10000)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6, 1, 10, 100, 10000)
	})
}
