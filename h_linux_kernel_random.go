package rng

import (
	"encoding/binary"
	"os"
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

func (k KernelRandom) Uint64() uint64 {
	var buf [8]byte
	h, err := os.Open("/dev/random")
	if err != nil {
		return 0
	}
	defer h.Close()
	if _, err := h.Read(buf[:]); err != nil {
		return 0
	}
	return binary.LittleEndian.Uint64(buf[:])
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
