package rng

// Galois LSFR (left shift) implementation.
type lsfrGL struct {
	seed uint64
}

func (r *lsfrGL) Seed(seed int64) {
	r.seed = uint64(seed)
}

func (r *lsfrGL) Int63() int64 {
	return int64(r.Uint64())
}

func (r *lsfrGL) Uint64() uint64 {
	lfsr := r.seed
	var period uint64
	initial := r.seed

	const tapMask uint16 = 0x002D

	for {
		msb := (lfsr >> 15) & 1
		lfsr <<= 1
		if msb == 1 {
			lfsr ^= uint64(tapMask)
		}
		period++
		if lfsr == initial {
			break
		}
	}
	return period
}
