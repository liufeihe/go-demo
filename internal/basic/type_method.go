package basic

import (
	"fmt"
	"time"
)

type T struct{}

func (t T) M(n int) {}

type field struct {
	name string
}

func (p field) print() {
	fmt.Println(p.name)
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
}
