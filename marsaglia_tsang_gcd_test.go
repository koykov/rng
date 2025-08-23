package rng

import (
	"math"
	"strconv"
	"testing"
)

func TestMarsagliaTsangGCD(t *testing.T) {
	gcd := func(a, b uint64) uint64 {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	testfn := func(rng Interface, numPairs int) (float64, float64, float64) {
		countCoprime := 0
		for i := 0; i < numPairs; i++ {
			a := rng.Uint64()
			b := rng.Uint64()
			if gcd(a, b) == 1 {
				countCoprime++
			}
		}
		observedProb := float64(countCoprime) / float64(numPairs)
		expectedProb := 6.0 / (math.Pi * math.Pi)
		stdDev := math.Sqrt(expectedProb * (1 - expectedProb) / float64(numPairs))
		zScore := (observedProb - expectedProb) / stdDev
		return observedProb, expectedProb, zScore
	}
	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				_, _, zScore := testfn(rng, step)
				if math.Abs(zScore) >= 3.0 {
					t.Fail()
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
