package rng

type xorshift64 struct {
	a uint64
}

func (r *xorshift64) Seed(_ int64) {}

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
