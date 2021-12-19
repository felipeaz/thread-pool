package polygons

import "math/rand"

type Triangle struct {
	Width  float64
	Height float64
	Length float64
}

func NewRandomTriangle(decimalRange float64) Triangle {
	return Triangle{
		Width:  rand.Float64() * decimalRange,
		Height: rand.Float64() * decimalRange,
		Length: rand.Float64() * decimalRange,
	}
}

func GenerateTriangleSlice(amount int, decimalRange float64) []Triangle {
	var triangles []Triangle
	for current := 0; current < amount; current++ {
		triangles = append(triangles, NewRandomTriangle(decimalRange))
	}
	return triangles
}

func (t Triangle) GetArea() float64 {
	return (t.Width * t.Height) / 2
}

func (t Triangle) GetVolume() float64 {
	return t.GetArea() * t.Length
}

func (t Triangle) GetPolygonType() string {
	return "Triangle"
}
