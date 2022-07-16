package basic

import "fmt"

func f() int64 {
	return 1
}

func testChannel() {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
		// go func(i, v int) {
		// 	fmt.Println(i, v)
		// }(i, v)
	}
	fmt.Println("end")
}

func testLoop() {
	// for
	var sl = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("sl[%d]=%d\n", i, sl[i])
	}
	for i, v := range sl {
		fmt.Printf("sl[%d]=%d\n", i, v)
	}
	var t1 = []int{0, 1, 2, 3, 4}
	var t2 = []int{10, 11, 12, 13, 14}
	fmt.Println("continue")
loop:
	for i := 0; i < len(t1); i++ {
		for j := 0; j < len(t2); j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if j == 2 {
				continue loop
			}
		}
	}
	fmt.Println("break")
loop2:
	for i := 0; i < len(t1); i++ {
		for j := 0; j < len(t2); j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if j == 2 {
				break loop2
			}
		}
	}
}

func testSwitch() {
	a := 1
	switch a {
	case 1, 2, 3, 4, 5:
		fmt.Println("it is a work day")
	case 6, 7:
		fmt.Println("it is a weekend day")
	default:
		fmt.Println("are you live on earth")
	}

	var x interface{} = 13
	switch v := x.(type) {
	case nil:
		fmt.Println("v is nil")
	case int:
		fmt.Println("v is int, value=", v)
	case string:
		fmt.Println("v is string, value=", v)
	case bool:
		fmt.Println("v is bool, value=", v)
	default:
		fmt.Println("dont't support the type", v)
	}
}

func CtrlDemo() {
	// if
	if a := f(); a > 0 {
		fmt.Println(a)
	}

	testLoop()

	testChannel()

	// switch
	testSwitch()
}
