package point

import "math"

type Point struct {
	x float64
	y float64
}
type Distance struct {
	Point
}

func NewPoints(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
func (d *Distance) DistanceBetweenPoints(p1 *Point, p2 *Point) float64 {
	dist := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	return dist
}
