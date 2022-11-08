package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	// As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()//io.Writer

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)

	//Kind returns the specific kind of this type.
	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan

	//Elem returns a type's element type. It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
	tm := tc.Elem()//one of Array, Chan, Map, Pointer, or Slice
	ta, tb := tm.Key(), tm.Elem()
	// The next line prints: map array slice
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind())
	tx, ty := ta.Elem(), tb.Elem()

	// byte is an alias of uint8
	fmt.Println(tx.Kind(), ty.Kind()) // int16 uint8
	fmt.Println(tx.Bits(), ty.Bits()) // 16 8
	fmt.Println(tx.ConvertibleTo(ty)) // true
	fmt.Println(tb.ConvertibleTo(ta)) // false

	// Slice and map types are incomparable.
	fmt.Println(tb.Comparable()) // false
	fmt.Println(tm.Comparable()) // false
	fmt.Println(ta.Comparable()) // true
	fmt.Println(tc.Comparable()) // true
}

/*
// Golang program to illustrate
// reflect.TypeOf() Function

package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {

	tst1 := "string"
	tst2 := 10
	tst3 := 1.2
	tst4 := true
	tst5 := []string{"foo", "bar", "baz"}
	tst6 := map[string]int{"apple": 23, "tomato": 13}

	
	// use of TypeOf method	
	fmt.Println(reflect.TypeOf(tst1)) //string
	fmt.Println(reflect.TypeOf(tst2)) //int
	fmt.Println(reflect.TypeOf(tst3)) //float64
	fmt.Println(reflect.TypeOf(tst4)) //bool
	fmt.Println(reflect.TypeOf(tst5)) //[]string
	fmt.Println(reflect.TypeOf(tst6)) //map[string]int

}

*/