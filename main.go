package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	firstSide  float64
	secondSide float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) getSquare() float64 {
	return r.secondSide * r.firstSide
}

func (c *Circle) getSquare() float64 {
	return c.radius * c.radius * math.Pi
}

type Shape interface {
	getSquare() float64
}

func main() {
	var one, two Shape
	rectangle := &Rectangle{firstSide: 12.0, secondSide: 11.2}
	circle := &Circle{radius: 22}
	one = rectangle
	two = circle
	fmt.Println(one.getSquare(), two.getSquare())

}
