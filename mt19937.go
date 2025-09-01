package rng

type mt19937 struct {
	seed uint64
}

func (r *mt19937) Seed(seed int64) {
	r.seed = uint64(seed)
}

func (r *mt19937) Int63() int64 {
	return int64(r.Uint64())
}

func (r *mt19937) Uint64() uint64 {
	// todo implement me
	return r.seed
}
