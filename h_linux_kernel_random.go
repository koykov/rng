package rng

import (
	"encoding/binary"
	"math"
	"os"
	"unsafe"
)

type KernelRandom struct{}

func (k KernelRandom) Seed(_ int64) {}

func (k KernelRandom) Int() int { return int(uint(k.Int63()) << 1 >> 1) }

func (k KernelRandom) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(k.Int31n(int32(n)))
	}
	return int(k.Int63n(int64(n)))
}

func (k KernelRandom) Int31() int32 { return int32(k.Int63() >> 32) }

func (k KernelRandom) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return k.Int31() & (n - 1)
	}
	mx := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := k.Int31()
	for v > mx {
		v = k.Int31()
	}
	return v % n
}

func (k KernelRandom) Int63() int64 { return int64(k.Uint64()) }

func (k KernelRandom) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return k.Int63() & (n - 1)
	}
	mx := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := k.Int63()
	for v > mx {
		v = k.Int63()
	}
	return v % n
}

func (k KernelRandom) Uint() uint { return uint(k.Uint64()) }

func (k KernelRandom) Uint32() uint32 { return uint32(k.Int63() >> 31) }

func (k KernelRandom) Uint64() (o uint64) {
	buf := *(*[]byte)(unsafe.Pointer(&o))
	if _, err := k.Read(buf); err != nil {
		return
	}
	return binary.LittleEndian.Uint64(buf[:]) // todo avoid LE
}

func (k KernelRandom) Float64() float64 {
again:
	f := float64(k.Int63()) / (1 << 63)
	if f == 1 {
		goto again
	}
	return f
}

func (k KernelRandom) Float32() float32 {
again:
	f := float32(k.Float64())
	if f == 1 {
		goto again
	}
	return f
}

func (k KernelRandom) ExpFloat64() float64 {
	for {
		j := k.Uint32()
		i := j & 0xFF
		x := float64(j) * float64(we[i])
		if j < ke[i] {
			return x
		}
		if i == 0 {
			return re - math.Log(k.Float64())
		}
		if fe[i]+float32(k.Float64())*(fe[i-1]-fe[i]) < float32(math.Exp(-x)) {
			return x
		}
	}
}

func (k KernelRandom) NormFloat64() float64 {
	for {
		j := int32(k.Uint32())
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
				x = -math.Log(k.Float64()) * (1.0 / rn)
				y := -math.Log(k.Float64())
				if y+y >= x*x {
					break
				}
			}
			if j > 0 {
				return rn + x
			}
			return -rn - x
		}
		if fn[i]+float32(k.Float64())*(fn[i-1]-fn[i]) < float32(math.Exp(-.5*x*x)) {
			return x
		}
	}
}

func (k KernelRandom) Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := k.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

func (k KernelRandom) Read(p []byte) (n int, err error) {
	h, err := os.Open("/dev/random")
	if err != nil {
		return
	}
	defer func() { _ = h.Close() }()
	return h.Read(p)
}
