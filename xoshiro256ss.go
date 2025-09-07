package rng

type xoshiro256ss struct {
	xoshiroBase
}

func newXoshiro256ss(seed uint64) *xoshiro256ss {
	r := &xoshiro256ss{xoshiroBase{s: [4]uint64{seed}}}
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
