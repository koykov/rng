package rng

// Fibonacci LSFR implementation.
type lsfrF struct {
	seed uint64
}

func (r *lsfrF) Seed(seed int64) {
	r.seed = uint64(seed)
}

func (r *lsfrF) Int63() int64 {
	return int64(r.Uint64())
}

func (r *lsfrF) Uint64() uint64 {
	lfsr := uint16(r.seed)
	var period uint64
	var fb uint16
	initial := lfsr
	for {
		fb = ((lfsr >> 0) ^ (lfsr >> 2) ^ (lfsr >> 3) ^ (lfsr >> 5)) & 1
		lfsr = (lfsr >> 1) | (fb << 15)
		period++
		if lfsr == initial {
			break
		}
	}
	r.seed = period
	return r.seed
}
