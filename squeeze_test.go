package rng

import (
	"math"
	"strconv"
	"testing"
)

func TestSqueeze(t *testing.T) {
	testfn := func(rng Interface, samples int, multiplier float64) (float64, []float64) {
		squeezed := make([]float64, samples)

		for i := 0; i < samples; i++ {
			original := math.Abs(rng.Float64())
			squeezed[i] = math.Mod(original*multiplier, 1.0)
		}

		bins := 10
		expected := float64(samples) / float64(bins)
		observed := make([]float64, bins)

		for _, val := range squeezed {
			bin := int(math.Floor(val * float64(bins)))
			if bin >= bins {
				bin = bins - 1
			}
			observed[bin]++
		}

		chi2 := 0.0
		for _, obs := range observed {
			chi2 += math.Pow(obs-expected, 2) / expected
		}

		return chi2, squeezed
	}
	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				multiplier := math.Pow(2, 32)
				chi2, _ := testfn(rng, step, multiplier)
				criticalValue := 16.919
				if chi2 >= criticalValue {
					t.Errorf("chi2 %f overflows critical values %f", chi2, criticalValue)
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1e6)
	})
}
