package rng

type lsfr struct {
	data, poly, mask int64
}

func (r *lsfr) Seed(seed int64) {
	r.data = seed
}

func (r *lsfr) Int63() int64 {
	// todo implement me
	return r.data
}

func (r *lsfr) Uint64() uint64 {
	return uint64(r.Int63())
}
