package basic

import "fmt"

func ConstantDemo() {
	type myInt int
	const n myInt = 13
	var a int = 5
	const n2 = 14
	// fmt.Println(a+n) // 报错
	fmt.Println(a + n2)

	const (
		java   = iota + 1
		js     // 2
		golang // 3
	)
	fmt.Println(golang)
}
