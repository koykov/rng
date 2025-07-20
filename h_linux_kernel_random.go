package rng

import (
	"encoding/binary"
	"os"
)

type KernelRandom struct{}

func (k KernelRandom) Seed(_ int64) {}

func (k KernelRandom) Int() int { return int(uint(k.Int63()) << 1 >> 1) }

func (k KernelRandom) Int31() int32 { return int32(k.Uint63()) }

func (k KernelRandom) Int63() int64 { return int64(k.Uint63()) }

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

func (k KernelRandom) Uint() uint { return uint(k.Uint63()) }

func (k KernelRandom) Uint32() float32 { return float32(k.Uint63()) }

func (k KernelRandom) Uint63() uint64 {
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
