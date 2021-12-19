package polygons

import (
	"math"
	"math/rand"
)

type Circle struct {
	Diameter float64
}

func NewRandomCircle(decimalRange float64) Circle {
	return Circle{
		Diameter: rand.Float64() * decimalRange,
	}
}

func GenerateCircleSlice(amount int, decimalRange float64) []Circle {
	var circles []Circle
	for current := 0; current < amount; current++ {
		circles = append(circles, NewRandomCircle(decimalRange))
	}
	return circles
}

func (c Circle) getRadius() float64 {
	return c.Diameter / 2
}

func (c Circle) GetArea() float64 {
	return math.Pi * math.Pow(c.getRadius(), 2)
}

func (c Circle) GetVolume() float64 {
	return math.Pi * math.Pow(c.getRadius(), 3)
}

func (c Circle) GetPolygonType() string {
	return "Circle"
}
