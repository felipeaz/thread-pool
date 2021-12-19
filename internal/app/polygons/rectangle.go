package polygons

import "math/rand"

type Rectangle struct {
	Width  float64
	Height float64
	Length float64
}

func NewRandomRectangle(decimalRange float64) Rectangle {
	return Rectangle{
		Width:  rand.Float64() * decimalRange,
		Height: rand.Float64() * decimalRange,
		Length: rand.Float64() * decimalRange,
	}
}

func GenerateRectangleSlice(amount int, decimalRange float64) []Rectangle {
	var retangles []Rectangle
	for current := 0; current < amount; current++ {
		retangles = append(retangles, NewRandomRectangle(decimalRange))
	}
	return retangles
}

func (r Rectangle) GetArea() float64 {
	return r.Height * r.Length
}

func (r Rectangle) GetVolume() float64 {
	return r.GetArea() * r.Length
}

func (r Rectangle) GetPolygonType() string {
	return "Rectangle"
}
