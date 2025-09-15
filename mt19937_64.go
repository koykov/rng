package rng

import "math/rand"

const (
	mt19937_64n       = 312
	mt19937_64m       = 156
	mt19937_64matrixA = 0xB5026F5AA96619E9
	mt19937_64umask   = 0xFFFFFFFF80000000
	mt19937_64lmask   = 0x7FFFFFFF
)

type mt19937_64 struct {
	seed   uint64
	state  [mt19937_64n]uint64
	statei int
	mag01  [2]uint64
}

func NewMt19937_64Source(seed int64) rand.Source64 {
	r := &mt19937_64{
		statei: mt19937_64n + 1,
		mag01:  [2]uint64{0, mt19937_64matrixA},
	}
	r.Seed(seed)
	return r
}

func (r *mt19937_64) Seed(seed int64) {
	r.seed = uint64(seed)
	r.state[0] = uint64(seed)
	for i := 1; i < mt19937_64n; i++ {
		r.state[i] = 6364136223846793005 * (r.state[i-1] ^ (r.state[i-1] >> 62) + uint64(i))
	}
	r.statei = mt19937_64n
}

func (r *mt19937_64) Int63() int64 {
	return int64(r.Uint64())
}

func (r *mt19937_64) Uint64() (x uint64) {
	var i int
	if r.statei >= mt19937_64n {
		if r.statei == mt19937_64n+1 {
			r.Seed(rand.Int63())
		}
		for i = 0; i < mt19937_64n-mt19937_64m; i++ {
			x = (r.state[i] & mt19937_64umask) | (r.state[i+1] & mt19937_64lmask)
			r.state[i] = r.state[i+mt19937_64m] ^ (x >> 1) ^ r.mag01[(int)(x&uint64(1))]
		}
		for ; i < mt19937_64n-1; i++ {
			x = (r.state[i] & mt19937_64umask) | (r.state[i+1] & mt19937_64lmask)
			r.state[i] = r.state[i+(mt19937_64m-mt19937_64n)] ^ (x >> 1) ^ r.mag01[(int)(x&uint64(1))]
		}
		x = (r.state[mt19937_64n-1] & mt19937_64umask) | (r.state[0] & mt19937_64lmask)
		r.state[mt19937_64n-1] = r.state[mt19937_64m-1] ^ (x >> 1) ^ r.mag01[(int)(x&uint64(1))]

		r.statei = 0
	}

	x = r.state[r.statei]
	r.statei++

	x ^= (x >> 29) & uint64(0x5555555555555555)
	x ^= (x << 17) & uint64(0x71D67FFFEDA60000)
	x ^= (x << 37) & uint64(0xFFF7EEE000000000)
	x ^= x >> 43
	return
}
