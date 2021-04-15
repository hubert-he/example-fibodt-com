package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

func describe(s shape) {
	fmt.Println("My area is ", s.Area())
}

type circle struct {
	radius float64
}

//This is how we define method on a struct
func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (cp *circle) SetRadius(value float64) {
	cp.radius = value
}

func main() {
	describe(circle{radius: 10})
	myc := circle{radius: 10}
	describe(myc)
	myc.SetRadius(30)
	describe(myc)
} //end
