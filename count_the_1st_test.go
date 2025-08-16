package rng

import (
	"math"
	"strconv"
	"testing"
)

func TestCountThe1st(t *testing.T) {
	onesCountN := func(x uint64, bpn int) (c int) {
		for i := 0; i < bpn; i++ {
			c += int((x >> i) & 1)
		}
		return
	}
	testfn := func(rng Interface, n, bpn int) (pValue, z float64) {
		totalBits := n * bpn
		expectedOnes := float64(totalBits) / 2.0
		var c int
		for i := 0; i < n; i++ {
			c += onesCountN(rng.Uint64(), bpn)
		}
		z = (float64(c) - expectedOnes) / math.Sqrt(expectedOnes*0.5)
		pValue = math.Erfc(math.Abs(z) / math.Sqrt(2))
		return
	}
	testgroup := func(t *testing.T, rng Interface, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) {
				pValue, _ := testfn(rng, step, 64)
				if pValue < 0.01 || pValue > 0.99 {
					t.Errorf("bad rate %f", pValue)
				}
			})
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		testgroup(t, KernelRandom, 1000, 1e4, 1e5, 1e6)
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		testgroup(t, KernelUrandom, 1000, 1e4, 1e5)
	})
}
