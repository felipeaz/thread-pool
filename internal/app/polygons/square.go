package polygons

import (
	"math"
	"math/rand"
)

type Square struct {
	Side float64
}

func NewRandomSquare(decimalRange float64) Square {
	return Square{
		Side: rand.Float64() * decimalRange,
	}
}

func GenerateRectangleSlices(amount int, decimalRange float64) []Square {
	var squares []Square
	for current := 0; current < amount; current++ {
		squares = append(squares, NewRandomSquare(decimalRange))
	}
	return squares
}

func (s Square) GetArea() float64 {
	return math.Pow(s.Side, 2)
}

func (s Square) GetVolume() float64 {
	return math.Pow(s.Side, 3)
}

func (s Square) GetPolygonType() string {
	return "Square"
}
