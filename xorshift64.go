package rng

type xorshift64 struct{}

func (r *xorshift64) Seed(seed int64) {}

func (r *xorshift64) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift64) Uint64() uint64 {
	// todo implement me
	return 0
}
