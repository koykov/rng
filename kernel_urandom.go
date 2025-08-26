package rng

import "math/rand"

const fpDevUrandom = "/dev/urandom"

var KernelUrandom = &kernelRandomWrapper{
	Rand:       rand.New(&kernelRandom{fp: fpDevUrandom}),
	Concurrent: &kernelRandomConcurrent{fp: fpDevUrandom},
}

func NewKernelUrandom() Interface {
	return &kernelRandomWrapper{
		Rand:       rand.New(&kernelRandom{fp: fpDevUrandom}),
		Concurrent: &kernelRandomConcurrent{fp: fpDevUrandom},
	}
}

var _ = NewKernelUrandom
