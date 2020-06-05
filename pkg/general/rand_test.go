package general

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIntRange_NextRandom(t *testing.T) {
	ir := IntRange{0,100*1024*1024, rand.New(rand.NewSource(42))}
	fmt.Println(ir.NextRandom())
	fmt.Println(ir.NextRandom())
	fmt.Println(ir.NextRandom())
}