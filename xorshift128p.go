package rng

type xorshift128p struct {
	x [2]uint64
}

func newXorshift128p(seed uint64) *xorshift128p {
	r := &xorshift128p{x: [2]uint64{seed}}
	return r
}

func (r *xorshift128p) Seed(seed int64) {
	r.x[0] = uint64(seed)
}

func (r *xorshift128p) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift128p) Uint64() uint64 {
	t := r.x[0]
	s := r.x[1]
	r.x[0] = s
	t ^= t << 23
	t ^= t >> 18
	t ^= s ^ (s >> 5)
	r.x[1] = t
	return t + s
}
