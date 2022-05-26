package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

var (
	errorOutOfCapacity   = errors.New("max capacity reached")
	errorIndexOutOfRange = errors.New("index out of range")
	errorNothingToRemove = errors.New("nothing to remove")
)

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errorOutOfCapacity
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errorIndexOutOfRange
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	extracted, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("nothing to extract: %w", err)
	}

	shapes := make([]Shape, 0, len(b.shapes)-1)
	for j := 0; j < i; j++ {
		shapes = append(shapes, b.shapes[i])
	}

	for j := i + 1; j < len(b.shapes); j++ {
		shapes = append(shapes, b.shapes[i])
	}

	b.shapes = shapes
	return extracted, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	removed, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("nothing to replace: %w", err)
	}

	shapes := make([]Shape, 0, len(b.shapes))
	for j := 0; j < i; j++ {
		shapes = append(shapes, b.shapes[i])
	}

	shapes = append(shapes, shape)
	for j := i + 1; j < len(b.shapes); j++ {
		shapes = append(shapes, b.shapes[i])
	}

	b.shapes = shapes
	return removed, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (result float64) {
	for i := range b.shapes {
		result += b.shapes[i].CalcPerimeter()
	}

	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (result float64) {
	for i := range b.shapes {
		result += b.shapes[i].CalcArea()
	}

	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	shapes := make([]Shape, 0)
	for i := range b.shapes {
		if _, ok := b.shapes[i].(Circle); ok {
			shapes = append(shapes, b.shapes[i])
		}
	}

	if len(b.shapes) == len(shapes) {
		return errorNothingToRemove
	}

	b.shapes = shapes
	return nil
}
