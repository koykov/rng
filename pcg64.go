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
	r.a = uint64(seed)
	r.b = uint64(seed)
	r.Uint64()
}

func (r *pcg64) Int63() int64 {
	return int64(r.Uint64())
}

func (r *pcg64) Uint64() uint64 {
	mul := func(a, b uint64) (a1, b1 uint64) {
		x0 := a & math.MaxUint32
		x1 := a >> 32
		y0 := b & math.MaxUint32
		y1 := b >> 32
		w0 := x0 * y0
		t := x1*y0 + w0>>32
		w1 := t & math.MaxUint32
		w2 := t >> 32
		w1 += x0 * y1
		a1 = x1*y1 + w2 + w1>>32
		b1 = a * b
		return
	}
	b, a := mul(r.a, pcg64multiplierA)
	b += r.b * pcg64multiplierA
	b += r.a * pcg64multiplierB
	r.a, r.b = a, b

	var c uint64
	r.a, c = r.add(r.a, pcg64incrementA, c)
	r.b, c = r.add(r.b, pcg64incrementB, c)

	return bits.RotateLeft64(r.b^r.a, -int(r.b>>58))
}

func (r *pcg64) add(a, b, c uint64) (uint64, uint64) {
	x := a + b + c
	return x, ((a & b) | (a|b)&^x) >> 63
}
