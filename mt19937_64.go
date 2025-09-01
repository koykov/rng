package rng

type mt19937_64 struct {
	seed uint64
}

func (r *mt19937_64) Seed(seed int64) {
	r.seed = uint64(seed)
}

func (r *mt19937_64) Int63() int64 {
	return int64(r.Uint64())
}

func (r *mt19937_64) Uint64() uint64 {
	// todo implement me
	return r.seed
}
