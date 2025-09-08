package rng

type xoshiro256p struct {
	xoshiroBase
}

func newXoshiro256p(state [4]uint64) *xoshiro256p {
	r := &xoshiro256p{xoshiroBase{s: state}}
	return r
}

func (r *xoshiro256p) Seed(seed int64) {
	r.s[0] = uint64(seed)
}

func (r *xoshiro256p) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xoshiro256p) Uint64() uint64 {
	result := r.s[0] + r.s[3]
	t := r.s[1] << 17

	r.s[2] ^= r.s[0]
	r.s[3] ^= r.s[1]
	r.s[1] ^= r.s[2]
	r.s[0] ^= r.s[3]

	r.s[2] ^= t
	r.s[3] = r.rol64(r.s[3], 45)

	return result
}
