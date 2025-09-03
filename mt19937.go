package rng

const (
	mt19937n     = 624
	mt19937m     = 397
	mt19937f     = 1812433253
	mt19937w     = 32
	mt19937r     = 31
	mt19937umask = uint64(0xffffffff << mt19937r)
	mt19937lmask = 0xffffffff >> (mt19937w - mt19937r)
	mt19937a     = 0x9908b0df
	mt19937u     = 11
	mt19937s     = 7
	mt19937t     = 15
	mt19937l     = 18
	mt19937b     = 0x9d2c5680
	mt19937c     = 0xefc60000
)

type mt19937 struct {
	seed   uint32
	state  [mt19937n]uint32
	statei int
}

func (r *mt19937) Seed(seed int64) {
	r.seed = uint32(seed)
	r.state[0] = uint32(seed)
	for i := 1; i < mt19937n; i++ {
		r.seed = mt19937f*(r.seed^(r.seed>>(mt19937w-2))) + 1
		r.state[i] = r.seed
	}
	r.statei = 0
}

func (r *mt19937) Int63() int64 {
	return int64(r.Uint64())
}

func (r *mt19937) Uint64() uint64 {
	k := r.statei
	j := k - (mt19937n - 1)
	for j < 0 {
		j += mt19937n
	}
	x := uint32(uint64(r.state[k])&mt19937umask) | (r.state[j] & mt19937lmask)
	xA := x >> 1
	if x&0x00000001 == 0 {
		xA ^= mt19937a
	}
	j = k - (mt19937n - mt19937m)
	if j < 0 {
		j += mt19937n
	}
	x = r.state[j] ^ xA
	r.state[k] = x
	k++
	if k >= mt19937n {
		k = 0
	}
	r.statei = k
	y := x ^ (x >> mt19937u)
	y = y ^ ((y << mt19937s) & mt19937b)
	y = y ^ ((y << mt19937t) & mt19937c)
	z := y ^ (y >> mt19937l)
	return uint64(z)
}
