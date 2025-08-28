package rng

import (
	"strconv"
	"sync"
	"testing"
)

func TestRNG(t *testing.T) {
	testfn := func(t *testing.T, rng Interface, n int, async bool) {
		if async {
			var wg sync.WaitGroup
			for i := 0; i < n; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if rng.Int63() == 0 {
						t.Error()
					}
				}()
			}
			wg.Wait()
			return
		}
		for i := 0; i < n; i++ {
			if rng.Int63() == 0 {
				t.Error()
			}
		}
	}
	testgroup := func(t *testing.T, rng Interface, async bool, steps ...int) {
		for _, step := range steps {
			t.Run(strconv.Itoa(step), func(t *testing.T) { testfn(t, rng, step, async) })
		}
	}
	t.Run("kernel/random", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, KernelRandom, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, KernelRandom.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, KernelUrandom, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, KernelUrandom.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/ZXSpectrum", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.ZXSpectrum, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.ZXSpectrum.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/ranqd1", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.Ranqd1, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.Ranqd1.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/Borland C++", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.BorlandCpp, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.BorlandCpp.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/Borland Delphi", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.BorlandDelphi, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.BorlandDelphi.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/Turbo Pascal", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.TurboPascal, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.TurboPascal.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/glibc", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.Glibc, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.Glibc.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/ANSI C", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.ANSI_C, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.ANSI_C.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/Microsoft Visual C++", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.MSVCpp, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.MSVCpp.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/Microsoft Visual Basic", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.MSVBasic, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.MSVBasic.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/RtlUniform", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.RtlUniform, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.RtlUniform.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/minstd_rand", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.MinstdRand, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.MinstdRand.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/minstd_rand0", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.MinstdRand0, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.MinstdRand0.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/MMIX", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.MMIX, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.MMIX.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/Musl", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.Musl, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.Musl.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/Java", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.Java, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.Java.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/POSIX", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.POSIX, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.POSIX.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/random0", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.Random0, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.Random0.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/cc65", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.Cc65, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.Cc65.Concurrent, true, 1, 10, 100, 1000)
		})
	})
	t.Run("lcg/RANDU", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, LCG.RANDU, false, 1, 10, 100, 1000)
		})
		t.Run("async", func(t *testing.T) {
			testgroup(t, LCG.RANDU.Concurrent, true, 1, 10, 100, 1000)
		})
	})
}

func BenchmarkRNG(b *testing.B) {
	benchfn := func(b *testing.B, rng Interface, async bool) {
		b.ReportAllocs()
		if async {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if rng.Int63() == 0 {
						b.Error()
					}
				}
			})
			return
		}
		for i := 0; i < b.N; i++ {
			_ = rng.Int63()
		}
	}
	b.Run("kernel/random", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, KernelRandom, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, KernelRandom.Concurrent, true) })
	})
	b.Run("kernel/urandom", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, KernelUrandom, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, KernelUrandom.Concurrent, true) })
	})
	b.Run("lcg/ZXSpectrum", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.ZXSpectrum, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.ZXSpectrum.Concurrent, true) })
	})
	b.Run("lcg/ranqd1", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.Ranqd1, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.Ranqd1.Concurrent, true) })
	})
	b.Run("lcg/Borland C++", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.BorlandCpp, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.BorlandCpp.Concurrent, true) })
	})
	b.Run("lcg/Borland Delphi", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.BorlandDelphi, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.BorlandDelphi.Concurrent, true) })
	})
	b.Run("lcg/Turbo Pascal", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.TurboPascal, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.TurboPascal.Concurrent, true) })
	})
	b.Run("lcg/glibc", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.Glibc, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.Glibc.Concurrent, true) })
	})
	b.Run("lcg/ANSI C", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.ANSI_C, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.ANSI_C.Concurrent, true) })
	})
	b.Run("lcg/Microsoft Visual C++", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.MSVCpp, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.MSVCpp.Concurrent, true) })
	})
	b.Run("lcg/Microsoft Visual Basic", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.MSVBasic, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.MSVBasic.Concurrent, true) })
	})
	b.Run("lcg/RtlUniform", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.RtlUniform, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.RtlUniform.Concurrent, true) })
	})
	b.Run("lcg/minstd_rand", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.MinstdRand, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.MinstdRand.Concurrent, true) })
	})
	b.Run("lcg/minstd_rand0", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.MinstdRand0, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.MinstdRand0.Concurrent, true) })
	})
	b.Run("lcg/MMIX", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.MMIX, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.MMIX.Concurrent, true) })
	})
	b.Run("lcg/Musl", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.Musl, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.Musl.Concurrent, true) })
	})
	b.Run("lcg/Java", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.Java, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.Java.Concurrent, true) })
	})
	b.Run("lcg/POSIX", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.POSIX, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.POSIX.Concurrent, true) })
	})
	b.Run("lcg/random0", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.Random0, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.Random0.Concurrent, true) })
	})
	b.Run("lcg/cc65", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.Cc65, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.Cc65.Concurrent, true) })
	})
	b.Run("lcg/RANDU", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) { benchfn(b, LCG.RANDU, false) })
		b.Run("async", func(b *testing.B) { benchfn(b, LCG.RANDU.Concurrent, true) })
	})
}
