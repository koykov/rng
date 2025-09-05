package rng

type xorshiftContainer struct {
	Xorshift32      wrapper
	Xorshift64      wrapper
	Xorshift128Plus wrapper
}
