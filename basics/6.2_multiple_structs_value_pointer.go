// ######### EXERCISE #########


























































package main

import (
	"fmt"
	"math"
)

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

func main() {
	rect := rectangle{
		width:  3,
		height: 4,
	}
	fmt.Println("Rectangle: ", rect)
	fmt.Println("Area: ", rect.area())
	fmt.Println("Perimeter: ", rect.perimeter())

	c := circle{
		radius: 3,
	}
	fmt.Println("\nCircle: ", c)
	fmt.Println("Area: ", c.area())
	fmt.Println("Perimeter: ", c.perimeter())
}
