package basic

import (
	"fmt"
	"unicode/utf8"
)

func StringDemo() {
	var c rune = 'a'
	fmt.Printf("char, %c \n", c)

	var str string = "中国人"
	// 字节视角
	fmt.Printf("%s, %d\n", str, len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("0x%x ", str[i])
	}
	fmt.Printf("\n")
	// 字符视角
	fmt.Printf("%s, %d\n", str, utf8.RuneCountInString(str))
	for _, c := range str {
		fmt.Printf("0x%x", c)
		fmt.Printf("(%c) ", c)
	}
	fmt.Printf("\n")
}
