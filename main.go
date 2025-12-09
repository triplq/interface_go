package main

import (
	"fmt"
	"strconv"
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
	return "Hello, Im " + p.name + " " + strconv.FormatUint(uint64(p.age), 10) + " y.o."
}

func (c Car) Describe() string {
	return c.color + " wroom-wroom " + strconv.FormatUint(uint64(c.speed), 10)
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
	PrintAll(car, pers)
}
