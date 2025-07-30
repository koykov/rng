package rng

import (
	"math"
	"strconv"
	"testing"
)

func TestUniformity(t *testing.T) {
	testfn := func(t *testing.T, rng Interface, n, bins int, min, max float64) {
		dist := make([]int, bins)
		for i := 0; i < n; i++ {
			r := math.Abs(rng.Float64())
			idx := int(math.Floor(r * float64(bins)))
			if idx >= bins {
				idx = bins - 1
			}
			dist[idx]++
		}

		exp := n / bins
		for i := 0; i < len(dist); i++ {
			if r := 100 * float64(dist[i]-exp) / float64(exp); r < min || r > max {
				t.Errorf("%d: %d - value out of min/max range %.2f%%", i, dist[i], 100*float64(dist[i]-exp)/float64(exp))
			}
		}
	}
	testgroup := func(t *testing.T, rng Interface, n int, min, max float64, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) { testfn(t, rng, n, step, min, max) })
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, KernelRandom, 1e6, -3.0, 3.0, 10, 20, 50, 100)
		})
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, KernelUrandom, 1e6, -3.0, 3.0, 10, 20, 50, 100)
		})
	})
}
