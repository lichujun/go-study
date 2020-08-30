package demo

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type Person struct {
	*Name
	Printer
	age int
}

func (p *Person) Print() {
	fmt.Println("hello")
}

func (p *Person) SaySome(some string) {
	fmt.Println(some)
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

func Test(t *testing.T) {
	p := Person{
		Name:    new(Name),
		age:     0,
		Printer: &PrintThing1{name: "joseph"},
	}
	fmt.Println(*p.Name)
	p.Print()

	fmt.Println(reflect.TypeOf(p).Field(2).Type)
}

var sc sync.Once

func GetOne() {
	sc.Do(func() {
		fmt.Println("just once")
	})
}

func TestGetOne(t *testing.T) {
	GetOne()
	GetOne()
	GetOne()
}
