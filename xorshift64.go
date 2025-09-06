package rng

type xorshift64 struct {
	a uint64
}

func newXorshift64(seed int64) *xorshift64 {
	r := &xorshift64{a: uint64(seed)}
	return r
}

func (r *xorshift64) Seed(seed int64) {
	r.a = uint64(seed)
}

func (r *xorshift64) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift64) Uint64() uint64 {
	x := r.a
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	r.a = x
	return x
}
