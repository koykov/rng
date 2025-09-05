package rng

type xorshift32 struct{}

func (r *xorshift32) Seed(seed int64) {}

func (r *xorshift32) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift32) Uint64() uint64 {
	// todo implement me
	return 0
}
