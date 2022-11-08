package main

import (
	"fmt"
	"reflect"
)

type Printer interface {
	Print()
	print2()
	print3()
}

type Document struct {
	Name string
}

type User struct {
	Name string
	Age  uint8
}

func (d *Document) Print() {
	fmt.Println("doc")
}

func (d *User) Hello() {
	fmt.Println("user hello")
}

// func getPrinter(o *Document) Printer {
// 	if(o==nil){
// 		return nil
// 	}
// 	return o
// }

func main() {
	//printStructRtypeMethods(rtype)
	t := reflect.TypeOf(&User{})
	
	for i := 0; i < t.NumMethod(); i++ {
	  m := t.Method(i)
	  fmt.Println(m.Name)
	}
}

func printStructMethods(any interface{}){
	t := reflect.TypeOf(any)
for i := 0; i < t.NumMethod(); i++ {
  m := t.Method(i)
  fmt.Println(m.Name)
}
}