package basic

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

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
	bs := []byte(str)
	fmt.Printf("byte: %x\n", bs)
	fmt.Printf("%s, %d\n", str, len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("0x%x ", str[i])
	}
	fmt.Printf("\n")
	fmt.Printf("[1]: %x, %c \n", str[1], str[1])

	// 字符视角
	rs := []rune(str)
	fmt.Printf("rune: %x\n", rs)
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

	// 字符串的底层实现
	var h = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&h))
	fmt.Printf("0x%x\n", hdr.Data)
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	dumpBytesArray((*p)[:])
}
