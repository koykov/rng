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
		Rand:       rand.New(NewLSFRFibonacciSource(rand.Uint64() % math.MaxUint16)),
		Concurrent: &Pool{New: func() rand.Source64 { return NewLSFRGaloisLeftShift(rand.Uint64() % math.MaxUint16) }},
	},
	GaloisLeftShift: wrapper{
		Rand:       rand.New(NewLSFRGaloisLeftShift(rand.Uint64() % math.MaxUint16)),
		Concurrent: &Pool{New: func() rand.Source64 { return NewLSFRGaloisLeftShift(rand.Uint64() % math.MaxUint16) }},
	},
	GaloisRightShift: wrapper{
		Rand:       rand.New(NewLSFRGaloisRightShift(rand.Uint64() % math.MaxUint16)),
		Concurrent: &Pool{New: func() rand.Source64 { return NewLSFRGaloisLeftShift(rand.Uint64() % math.MaxUint16) }},
	},
}
