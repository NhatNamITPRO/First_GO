package main

import "fmt"

type Move interface {
	move() string
}
type Car struct {
	name string
}
type Animal struct {
	name string
}

func (c Car) move() string {
	return "Moving by wheel"
}
func (a Animal) move() string {
	return "Moving by leg"
}
func main() {
	objs := []Move{Car{"BMW"},
		Animal{"dog"}}
	for _, obj := range objs {
		fmt.Println(obj.move())
	}
}
