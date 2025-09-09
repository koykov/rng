package rng

const (
	pcg32increment  = 1442695040888963407
	pcg32multiplier = 6364136223846793005
)

type pcg32 struct {
	pcgBase
}

func newPCG32(seed uint64) *pcg32 {
	var r pcg32
	r.Seed(int64(seed))
	return &r
}

func (r *pcg32) Seed(seed int64) {
	r.state = uint64(seed) + pcg32increment
	r.Uint64()
}

func (r *pcg32) Int63() int64 {
	return int64(r.Uint64())
}

func (r *pcg32) Uint64() uint64 {
	x := r.state
	count := x >> 59

	r.state = x*pcg32multiplier + pcg32increment
	x ^= x >> 18
	return uint64(r.rotr32(uint32(x>>27), count))
}
