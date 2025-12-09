package main

import (
	"fmt"
)

type Describable interface {
	Describe() string
}

type Person struct {
	name string
	age  uint32
}

type Car struct {
	color string
	speed uint32
}

func (p Person) Describe() string {
	return fmt.Sprintf("Hello, Im %s %d y.o.", p.name, p.age)
}

func (c Car) Describe() string {
	return fmt.Sprintf("%s wroom-wroom %d", c.color, c.speed)
}

func PrintAll(descs ...Describable) {
	for _, val := range descs {
		fmt.Println(val.Describe())
	}
}

func Print(d Describable) {
	fmt.Println(d.Describe())
}

func main() {
	car := Car{"Red", 120}
	pers := Person{"Alex", 20}
	Print(car)
	Print(pers)
	fmt.Println()
	PrintAll(car, pers)
}
