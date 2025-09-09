package rng

import (
	"math"
	"math/bits"
)

const (
	pcg64multiplier  = 47026247687942121848144207491837523525
	pcg64multiplierA = pcg64multiplier & math.MaxUint64
	pcg64multiplierB = pcg64multiplier >> 64

	pcg64increment  = 117397592171526113268558934119004209487
	pcg64incrementA = pcg64increment & math.MaxUint64
	pcg64incrementB = pcg64increment >> 64
)

type pcg64 struct {
	a, b uint64
}

func newPCG64(a, b uint64) *pcg64 {
	return &pcg64{a, b}
}

func (r *pcg64) Seed(seed int64) {
	r.a = uint64(seed) + pcg64increment
	r.b = uint64(seed) * pcg64multiplier
	r.Uint64()
}

func (r *pcg64) Int63() int64 {
	return int64(r.Uint64())
}

func (r *pcg64) Uint64() uint64 {
	// r.mul() todo implement me
	// r.add() todo implement me
	return bits.RotateLeft64(r.b^r.a, -int(r.b>>58))
}
