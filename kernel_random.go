package rng

import (
	"encoding/binary"
	"math"
	"os"
	"sync"
)

type kernelRandom struct {
	f    *os.File
	err  error
	once sync.Once
}

type kernelRandomWrapper struct {
	kernelRandom
	Concurrent kernelRandomConcurrent
}

var KernelRandom = &kernelRandomWrapper{}

func NewKernelRandom() Interface {
	return &kernelRandom{}
}

func (r *kernelRandom) Seed(_ int64) {}

func (r *kernelRandom) Int() int { return int(uint(r.Int63()) << 1 >> 1) }

func (r *kernelRandom) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

func (r *kernelRandom) Int31() int32 { return int32(r.Int63() >> 32) }

func (r *kernelRandom) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	mx := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > mx {
		v = r.Int31()
	}
	return v % n
}

func (r *kernelRandom) Int63() int64 { return int64(r.Uint64()) }

func (r *kernelRandom) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int63() & (n - 1)
	}
	mx := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > mx {
		v = r.Int63()
	}
	return v % n
}

func (r *kernelRandom) Uint() uint { return uint(r.Uint64()) }

func (r *kernelRandom) Uint32() uint32 { return uint32(r.Int63() >> 31) }

func (r *kernelRandom) Uint64() uint64 {
	var buf [8]byte
	if _, err := r.Read(buf[:]); err != nil {
		return 0
	}
	return binary.LittleEndian.Uint64(buf[:])
}

func (r *kernelRandom) Float64() float64 {
again:
	f := float64(r.Int63()) / (1 << 63)
	if f == 1 {
		goto again
	}
	return f
}

func (r *kernelRandom) Float32() float32 {
again:
	f := float32(r.Float64())
	if f == 1 {
		goto again
	}
	return f
}

func (r *kernelRandom) ExpFloat64() float64 {
	for {
		j := r.Uint32()
		i := j & 0xFF
		x := float64(j) * float64(we[i])
		if j < ke[i] {
			return x
		}
		if i == 0 {
			return re - math.Log(r.Float64())
		}
		if fe[i]+float32(r.Float64())*(fe[i-1]-fe[i]) < float32(math.Exp(-x)) {
			return x
		}
	}
}

func (r *kernelRandom) NormFloat64() float64 {
	for {
		j := int32(r.Uint32())
		i := j & 0x7F
		x := float64(j) * float64(wn[i])
		abs := uint32(j)
		if j < 0 {
			abs = uint32(-j)
		}
		if abs < kn[i] {
			// This case should be hit better than 99% of the time.
			return x
		}

		if i == 0 {
			// This extra work is only required for the base strip.
			for {
				x = -math.Log(r.Float64()) * (1.0 / rn)
				y := -math.Log(r.Float64())
				if y+y >= x*x {
					break
				}
			}
			if j > 0 {
				return rn + x
			}
			return -rn - x
		}
		if fn[i]+float32(r.Float64())*(fn[i-1]-fn[i]) < float32(math.Exp(-.5*x*x)) {
			return x
		}
	}
}

func (r *kernelRandom) Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

func (r *kernelRandom) Read(p []byte) (n int, err error) {
	if r.once.Do(r.init); r.err != nil {
		return 0, r.err
	}
	return r.f.Read(p)
}

func (r *kernelRandom) Shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}

	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.int31n(int32(i + 1)))
		swap(i, j)
	}
}

func (r *kernelRandom) int31n(n int32) int32 {
	v := r.Uint32()
	prod := uint64(v) * uint64(n)
	low := uint32(prod)
	if low < uint32(n) {
		thresh := uint32(-n) % uint32(n)
		for low < thresh {
			v = r.Uint32()
			prod = uint64(v) * uint64(n)
			low = uint32(prod)
		}
	}
	return int32(prod >> 32)
}

func (r *kernelRandom) init() {
	r.f, r.err = os.Open("/dev/random")
}

var _ = NewKernelRandom
