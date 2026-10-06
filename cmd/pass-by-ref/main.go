package main

import "fmt"

type cars struct {
	speed int
	name  string
	color string
}

func (c *cars) updateCars(speed int, color string, name string) {
	if speed != 0 {
		c.speed = speed
	}
	if color != "" {
		c.color = color
	}
	if name != "" {
		c.name = name
	}

}

func main() {
	myOwn := cars{speed: 10, name: "daio", color: "cherry"}
	myOwn.updateCars(200, "blue", "Ferrari")
	fmt.Println(myOwn)
}
