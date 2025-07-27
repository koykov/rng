package rng

const fpDevUrandom = "/dev/urandom"

var KernelUrandom = NewKernelUrandom().(*kernelRandomWrapper)

func NewKernelUrandom() Interface {
	return &kernelRandomWrapper{
		kernelRandom: kernelRandom{fp: fpDevUrandom},
		Concurrent:   &kernelRandomConcurrent{fp: fpDevUrandom},
	}
}
