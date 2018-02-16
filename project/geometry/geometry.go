package geometry

import (
	"fmt"
	"math"
)

// Define type geometry interface
type geometry interface {
	area() float64
	perimeter() float64
}

// Define type struct of rectangle
type rectangle struct {
	width  float64
	height float64
}

// Define type struct of circle
type circle struct {
	radius float64
}

// function which gives area of rectangle
func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

// function which gives perimter of rectangle
func (rect rectangle) perimeter() float64 {
	return 2*rect.width + 2*rect.height
}

// function which gives area of rectangle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// function which gives perimter of rectangle
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Print("\nGeometry: ", g)
	fmt.Print("\nArea: ", g.area())
	fmt.Print("\nPerimter: ", g.perimeter(), "\n")
}
