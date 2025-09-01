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
	const tapMask uint16 = 0x002D

	lsfr := uint16(r.seed)
	initial := uint16(r.seed)
	var period uint64
	for {
		msb := int16(lsfr) < 0
		lsfr <<= 1
		if msb {
			lsfr ^= tapMask
		}
		period++
		if lsfr == initial {
			break
		}
	}
	r.seed = period
	return r.seed
}

// Galois LSFR (right shift) implementation.
type lsfrGR struct {
	seed uint64
}

func (r *lsfrGR) Seed(seed int64) {
	r.seed = uint64(seed)
}

func (r *lsfrGR) Int63() int64 {
	return int64(r.Uint64())
}

func (r *lsfrGR) Uint64() uint64 {
	const tapMask uint16 = 0xB400

	lfsr := r.seed
	var period uint64
	initial := r.seed
	for {
		lsb := lfsr&1 == 1
		lfsr >>= 1
		if lsb {
			lfsr ^= uint64(tapMask)
		}
		period++
		if lfsr == initial {
			break
		}
	}
	return period
}
