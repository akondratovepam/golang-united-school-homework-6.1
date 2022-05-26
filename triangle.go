package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (triangle Triangle) CalcArea() float64 {
	s := 3 * triangle.Side / 2
	return math.Sqrt(s * 3 * (s - triangle.Side))
}

func (triangle Triangle) CalcPerimeter() float64 {
	return 3 * triangle.Side
}
