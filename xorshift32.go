package rng

type xorshift32 struct {
	a uint32
}

func (r *xorshift32) Seed(_ int64) {}

func (r *xorshift32) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift32) Uint64() uint64 {
	x := r.a
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.a = x
	return uint64(x)
}
