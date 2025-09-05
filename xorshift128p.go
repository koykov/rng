package rng

type xorshift128p struct{}

func (r *xorshift128p) Seed(seed int64) {}

func (r *xorshift128p) Int63() int64 {
	return int64(r.Uint64())
}

func (r *xorshift128p) Uint64() uint64 {
	// todo implement me
	return 0
}
