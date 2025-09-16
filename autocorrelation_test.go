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

	t.Run("lcg/ZXSpectrum", func(t *testing.T) {
		testgroup(t, LCG.ZXSpectrum, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/ranqd1", func(t *testing.T) {
		testgroup(t, LCG.Ranqd1, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/Borland C++", func(t *testing.T) {
		testgroup(t, LCG.BorlandCpp, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/Borland Delphi", func(t *testing.T) {
		testgroup(t, LCG.BorlandDelphi, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/Turbo Pascal", func(t *testing.T) {
		testgroup(t, LCG.TurboPascal, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/glibc", func(t *testing.T) {
		testgroup(t, LCG.Glibc, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/ANSI C", func(t *testing.T) {
		testgroup(t, LCG.ANSI_C, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/Microsoft Visual C++", func(t *testing.T) {
		testgroup(t, LCG.MSVCpp, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/Microsoft Visual Basic", func(t *testing.T) {
		testgroup(t, LCG.MSVBasic, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/RtlUniform", func(t *testing.T) {
		testgroup(t, LCG.RtlUniform, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/minstd_rand", func(t *testing.T) {
		testgroup(t, LCG.MinstdRand, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/minstd_rand0", func(t *testing.T) {
		testgroup(t, LCG.MinstdRand0, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/MMIX", func(t *testing.T) {
		testgroup(t, LCG.MMIX, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/Musl", func(t *testing.T) {
		testgroup(t, LCG.Musl, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/Java", func(t *testing.T) {
		testgroup(t, LCG.Java, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/POSIX", func(t *testing.T) {
		testgroup(t, LCG.POSIX, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/random0", func(t *testing.T) {
		testgroup(t, LCG.Random0, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/cc65", func(t *testing.T) {
		testgroup(t, LCG.Cc65, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lcg/RANDU", func(t *testing.T) {
		testgroup(t, LCG.RANDU, 1e6, 1, 10, 100, 1000, 10000)
	})

	t.Run("lsfr/Fibonacci", func(t *testing.T) {
		testgroup(t, LSFR.Fibonacci, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lsfr/GaloisLeftShift", func(t *testing.T) {
		testgroup(t, LSFR.GaloisLeftShift, 1e6, 1, 10, 100, 1000, 10000)
	})
	t.Run("lsfr/GaloisRightShift", func(t *testing.T) {
		testgroup(t, LSFR.GaloisRightShift, 1e6, 1, 10, 100, 1000, 10000)
	})

	t.Run("Mersenne Twister", func(t *testing.T) {
		t.Run("32", func(t *testing.T) {
			testgroup(t, MersenneTwister.mt19937, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("64", func(t *testing.T) {
			testgroup(t, MersenneTwister.mt19937_64, 1e6, 1, 10, 100, 1000, 10000)
		})
	})

	t.Run("Xorshift", func(t *testing.T) {
		t.Run("32", func(t *testing.T) {
			testgroup(t, Xorshift.Xorshift32, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("64", func(t *testing.T) {
			testgroup(t, Xorshift.Xorshift64, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("128", func(t *testing.T) {
			testgroup(t, Xorshift.Xorshift128, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("128p", func(t *testing.T) {
			testgroup(t, Xorshift.Xorshift128Plus, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("1024s", func(t *testing.T) {
			testgroup(t, Xorshift.Xorshift1024s, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("r128p", func(t *testing.T) {
			testgroup(t, Xorshift.Xorshiftr128Plus, 1e6, 1, 10, 100, 1000, 10000)
		})
	})

	t.Run("Xoshiro", func(t *testing.T) {
		t.Run("256p", func(t *testing.T) {
			testgroup(t, Xoshiro.Xoshiro256Plus, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("256pp", func(t *testing.T) {
			testgroup(t, Xoshiro.Xoshiro256PlusPlus, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("256ss", func(t *testing.T) {
			testgroup(t, Xoshiro.Xoshiro256SS, 1e6, 1, 10, 100, 1000, 10000)
		})
	})

	t.Run("PCG", func(t *testing.T) {
		t.Run("32", func(t *testing.T) {
			testgroup(t, PCG.PCG32, 1e6, 1, 10, 100, 1000, 10000)
		})
		t.Run("64", func(t *testing.T) {
			testgroup(t, PCG.PCG64, 1e6, 1, 10, 100, 1000, 10000)
		})
	})
}
