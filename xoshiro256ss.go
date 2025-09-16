package rng

import "math/rand"

type xoshiro256ss struct {
	xoshiroBase
}

func NewXoshiro256ssSource(state [4]uint64) rand.Source64 {
	r := &xoshiro256ss{xoshiroBase{s: state}}
	return r
}

func (r *xoshiro256ss) Seed(seed int64) {
	r.s[0] = uint64(seed)
}

func (r *xoshiro256ss) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xoshiro256ss) Uint64() uint64 {
	result := r.rol64(r.s[1]*5, 7) * 9
	t := r.s[1] << 17

	r.s[2] ^= r.s[0]
	r.s[3] ^= r.s[1]
	r.s[1] ^= r.s[2]
	r.s[0] ^= r.s[3]

	r.s[2] ^= t
	r.s[3] = r.rol64(r.s[3], 45)

	return result
}
