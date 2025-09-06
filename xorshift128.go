package rng

type xorshift128 struct {
	x [4]uint32
}

func newXorshift128(seed uint64) *xorshift128 {
	r := &xorshift128{x: [4]uint32{uint32(seed)}}
	return r
}

func (r *xorshift128) Seed(seed int64) {
	r.x[0] = uint32(seed)
}

func (r *xorshift128) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift128) Uint64() uint64 {
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
