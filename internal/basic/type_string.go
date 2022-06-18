package basic

import (
	"fmt"
	"unicode/utf8"
)

func encodeRune(r rune) {
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)
	fmt.Printf("Unicode Char: %c, utf8: %x \n", r, buf)
}

func decodeRune(buf []byte) {
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("utf8: %x, Unicode Char: %s \n", buf, string(r))
}

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
	fmt.Printf("[1]: %x, %c \n", str[1], str[1])

	// 字符视角
	fmt.Printf("%s, %d\n", str, utf8.RuneCountInString(str))
	for _, c := range str {
		fmt.Printf("0x%x", c)
		fmt.Printf("(%b)(%c) ", c, c)
	}
	fmt.Printf("\n")

	buf := []byte{0xe4, 0xb8, 0xad}
	decodeRune(buf)
	encodeRune('a')
	encodeRune('b')
	encodeRune('中')
}
