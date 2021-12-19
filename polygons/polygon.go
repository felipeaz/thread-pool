package polygons

type Polygon interface {
	GetArea() float64
	GetVolume() float64
	GetPolygonType() string
}
