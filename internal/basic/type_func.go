package basic

import "fmt"

func testMultiReturn(flag int) (a int, b bool) {
	a = 1
	b = true
	if flag == 1 {
		b = false
		return
	}
	return
}

type newTestMultiReturn func(flag int) (a int, b bool)

func foo() {
	fmt.Println("call foo")
	bar()
	fmt.Println("exit foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic:", e)
		}
	}()
	fmt.Println("call bar")
	panic("panic occur in bar")
	zoo()
	fmt.Println("exit bar")
}

func zoo() {
	fmt.Println("call zoo")
	fmt.Println("exit zoo")
}

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func FuncDemo() {
	fmt.Println(testMultiReturn(0))
	var f1 newTestMultiReturn = testMultiReturn
	fmt.Println(f1(1))

	foo()

	fmt.Println("foo1 result:")
	foo1()
	fmt.Println("foo2 result:")
	foo2()
	fmt.Println("foo3 result:")
	foo3()
}
