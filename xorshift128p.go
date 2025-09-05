package rng

type xorshift128p struct {
	x [4]uint32
}

func (r *xorshift128p) Seed(_ int64) {}

func (r *xorshift128p) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift128p) Uint64() uint64 {
	t := r.x[3]
	s := r.x[0]
	r.x[3] = r.x[2]
	r.x[2] = r.x[1]
	r.x[1] = s

	t ^= t << 11
	t ^= t >> 8
	r.x[0] = t ^ s ^ (s >> 19)
	return uint64(r.x[0])
}
