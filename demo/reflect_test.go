package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectMethod(t *testing.T) {
	p := Person{
		Name:    &Name{name: "joseph"},
		Printer: &PrintThing{name: "joe"},
		age:     0,
	}
	p1 := reflect.ValueOf(p).Interface().(Person)
	p1.Print()

	pv := reflect.ValueOf(&p)

	printerType := reflect.TypeOf((*Printer)(nil)).Elem()

	fmt.Println(printerType)

	fmt.Println(reflect.ValueOf(p).FieldByName("Printer").Type().Implements(printerType))

	fmt.Println(pv.NumMethod())

	printMethod := pv.MethodByName("SaySome")

	var params = make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("say some")

	printMethod.Call(params)
}

func TestScanPackage(t *testing.T) {

}
