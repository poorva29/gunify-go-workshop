package main

import (
	geometry "project/geometry"
)

func main() {
	rect := geometry.rectangle{
		width:  3,
		height: 4,
	}
	geometry.measure(rect)

	c := geometry.circle{
		radius: 3,
	}
	geometry.measure(c)
}
