package main

import "fmt"

type Person struct {
	*Name
	Printer
	age int
}

type Name struct {
	name string
}

type Printer interface {
	Print()
}

type PrintThing struct {
	name string
}

type PrintThing1 struct {
	name string
}

func (p *PrintThing) Print() {
	fmt.Println(p.name + "0")
}

func (p *PrintThing1) Print() {
	fmt.Println(p.name + "1")
}

func main() {
	p := Person{
		Name:    &Name{name: "joseph"},
		age:     0,
		Printer: &PrintThing1{name: "joseph"},
	}
	fmt.Println(*p.Name)
	p.Print()
}
