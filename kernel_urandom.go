package rng

const fpDevUrandom = "/dev/urandom"

var KernelUrandom = &kernelRandomWrapper{
	kernelRandom: kernelRandom{fp: fpDevUrandom},
	Concurrent:   &kernelRandomConcurrent{fp: fpDevUrandom},
}

func NewKernelUrandom() Interface {
	return &kernelRandomWrapper{
		kernelRandom: kernelRandom{fp: fpDevUrandom},
		Concurrent:   &kernelRandomConcurrent{fp: fpDevUrandom},
	}
}

var _ = NewKernelUrandom
