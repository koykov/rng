package rng

import (
	"encoding/binary"
	"os"
)

type KernelRandom struct{}

func (k KernelRandom) Seed(_ int64) {}

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
