package main

import "fmt"

type Shape interface {
	Area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length, breadth, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	c1 := circle{radius: 33}
	r1 := rectangle{length: 10, breadth: 4, height: 3}

	c1Area := c1.Area()
	r1Area := r1.Area()

	fmt.Printf("Area of the circle is %.2f\n\n", c1Area)
	fmt.Printf("Area of the rectangle %.2f\n\n", r1Area)
}
