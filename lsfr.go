package rng

import (
	"math"
	"math/rand"
	"sync"
)

type lsfrContainer struct {
	Fibonacci        wrapper
	GaloisLeftShift  wrapper
	GaloisRightShift wrapper
}

var LSFR = &lsfrContainer{
	Fibonacci: wrapper{
		Rand:       rand.New(&lsfrF{seed: rand.Uint64() % math.MaxUint16}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lsfrF{seed: rand.Uint64() % math.MaxUint16}) }}},
	},
	GaloisLeftShift: wrapper{
		Rand:       rand.New(&lsfrGL{seed: rand.Uint64() % math.MaxUint16}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lsfrGL{seed: rand.Uint64() % math.MaxUint16}) }}},
	},
	GaloisRightShift: wrapper{
		Rand:       rand.New(&lsfrGR{seed: rand.Uint64() % math.MaxUint16}),
		Concurrent: &concurrent{Pool: sync.Pool{New: func() any { return rand.New(&lsfrGR{seed: rand.Uint64() % math.MaxUint16}) }}},
	},
}
