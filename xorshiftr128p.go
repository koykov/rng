package rng

import "math/rand"

type xorshiftr128p struct {
	s [2]uint64
}

func NewXorshiftr128pSource(seed uint64) rand.Source64 {
	r := &xorshiftr128p{s: [2]uint64{seed}}
	return r
}

func (r *xorshiftr128p) Seed(seed int64) {
	r.s[0] = uint64(seed)
}

func (r *xorshiftr128p) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshiftr128p) Uint64() uint64 {
	x := r.s[0]
	y := r.s[1]
	r.s[0] = y
	x ^= x << 23
	x ^= x >> 17
	x ^= y
	r.s[1] = x + y
	return x
}
