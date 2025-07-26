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
			testgroup(t, KernelRandom, true, 1, 10, 100, 1000)
		})
	})
	t.Run("kernel/urandom", func(t *testing.T) {
		t.Run("sync", func(t *testing.T) {
			testgroup(t, KernelUrandom, false, 1, 10, 100, 1000)
		})
		// t.Run("async", func(t *testing.T) {
		// 	testgroup(t, KernelUrandom, true, 1, 10, 100, 1000)
		// })
	})
}

func BenchmarkRNG(b *testing.B) {
	b.Run("kernel/random", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = KernelRandom.Int63()
			}
		})
		b.Run("async", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if KernelRandom.Int63() == 0 {
						b.Error()
					}
				}
			})
		})
	})
	b.Run("kernel/urandom", func(b *testing.B) {
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = KernelUrandom.Int63()
			}
		})
		// b.Run("async", func(b *testing.B) {
		// 	b.ReportAllocs()
		// 	b.RunParallel(func(pb *testing.PB) {
		// 		for pb.Next() {
		// 			if KernelUrandom.Int63() == 0 {
		// 				b.Error()
		// 			}
		// 		}
		// 	})
		// })
	})
}
