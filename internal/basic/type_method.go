package basic

import (
	"fmt"
	"reflect"
	"time"
)

type T struct{}

func (t T) M(n int)   {}
func (t *T) M1(n int) {}

type field struct {
	name string
}

func (p field) print() {
	fmt.Println(p.name)
}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)
	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}
	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty.\n", dynTyp)
		return
	}
	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

func MethodDemo() {
	var t T
	t.M(1)
	p := &T{}
	p.M(2)

	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}
	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print()
	}
	time.Sleep(3 * time.Second)

	dumpMethodSet(t)
	dumpMethodSet(&t)
}
