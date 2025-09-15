package rng

import (
	"math"
	"math/rand"
)

type lsfrContainer struct {
	Fibonacci        wrapper
	GaloisLeftShift  wrapper
	GaloisRightShift wrapper
}

var LSFR = &lsfrContainer{
	Fibonacci: wrapper{
		Rand:       rand.New(&lsfrF{seed: rand.Uint64() % math.MaxUint16}),
		Concurrent: &Pool{New: func() rand.Source64 { return &lsfrF{seed: rand.Uint64() % math.MaxUint16} }},
	},
	GaloisLeftShift: wrapper{
		Rand:       rand.New(&lsfrGL{seed: rand.Uint64() % math.MaxUint16}),
		Concurrent: &Pool{New: func() rand.Source64 { return &lsfrGL{seed: rand.Uint64() % math.MaxUint16} }},
	},
	GaloisRightShift: wrapper{
		Rand:       rand.New(&lsfrGR{seed: rand.Uint64() % math.MaxUint16}),
		Concurrent: &Pool{New: func() rand.Source64 { return &lsfrGR{seed: rand.Uint64() % math.MaxUint16} }},
	},
}
