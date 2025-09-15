package rng

import "math/rand"

type wrapper struct {
	*rand.Rand
	Concurrent *Pool
}
