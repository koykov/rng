package rng

type xorshift128 struct{}

func (r *xorshift128) Seed(seed int64) {}

func (r *xorshift128) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift128) Uint64() uint64 {
	// todo implement me
	return 0
}
