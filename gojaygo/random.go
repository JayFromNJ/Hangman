package gojaygo

import (
	"math/rand"
	"time"
)

type Random struct {
	rand *rand.Rand
}

func NewRandomGenerator() *Random {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	ra1 := Random{
		rand: r1,
	}

	return &ra1
}

// NextInt max inclusive
func (r *Random) NextInt(min int, max int) int {
	return r.rand.Intn(max+1) + min
}

func (r *Random) NextFloat32(min float32, max float32) float32 {
	return r.rand.Float32() * (max - min) + min
}

func (r *Random) NextFloat64(min float64, max float64) float64 {
	return r.rand.Float64() * (max - min) + min
}

func (r *Random) DiceRoll(sides int) int {
	return r.NextInt(1, sides)
}

func (r *Random) DiceRollN(sides int, times int) int {
	result := 0
	for i := 0; i < times; i++ {
		result = result + r.DiceRoll(sides)
	}
	return result
}

func (r *Random) CoinFlip() int {
	return r.NextInt(0, 1)
}