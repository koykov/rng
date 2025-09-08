package rng

type xoshiro256pp struct {
	xoshiroBase
}

func newXoshiro256pp(state [4]uint64) *xoshiro256pp {
	r := &xoshiro256pp{xoshiroBase{s: state}}
	return r
}

func (r *xoshiro256pp) Seed(seed int64) {
	r.s[0] = uint64(seed)
}

func (r *xoshiro256pp) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xoshiro256pp) Uint64() uint64 {
	result := r.rol64(r.s[0]+r.s[3], 23) + r.s[0]
	t := r.s[1] << 17

	r.s[2] ^= r.s[0]
	r.s[3] ^= r.s[1]
	r.s[1] ^= r.s[2]
	r.s[0] ^= r.s[3]

	r.s[2] ^= t
	r.s[3] = r.rol64(r.s[3], 45)

	return result
}
