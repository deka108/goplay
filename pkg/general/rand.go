package general

import "math/rand"

type IntRange struct {
	Min, Max int
	Rand *rand.Rand
}

func (ir *IntRange) NextRandom() int {
	return ir.Rand.Intn(ir.Max - ir.Min +1) + ir.Min
}
