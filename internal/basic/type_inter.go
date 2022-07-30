package basic

import "fmt"

type MyInterface interface {
	M1()
}

type MyT int

func (MyT) M1() {
	fmt.Println("m1")
}

func typeCheck() {
	var t MyT
	var i interface{} = 1
	v1, ok := i.(MyInterface)
	if !ok {
		fmt.Printf("%T \n", v1)
	}
	// v1 = 1
	var i2 interface{} = t
	v2, ok := i2.(MyInterface)
	if ok {
		fmt.Printf("%T \n", v2)
		fmt.Println(v2)
	}
}

func InterDemo() {
	typeCheck()
}
