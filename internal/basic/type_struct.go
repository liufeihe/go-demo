package basic

import (
	"fmt"
	"unsafe"
)

type Persion struct {
	Name  string
	Phone string
	Addr  string
}

type Book struct {
	Title  string
	Author Persion
}

type S1 struct {
	b byte
	i int64
	u uint16
}

type S2 struct {
	b byte
	u uint16
	i int64
}

func newBook(name string, phone string, addr string, title string) Book {
	b := Book{
		Title: title,
		Author: Persion{
			Name:  name,
			Phone: phone,
			Addr:  addr,
		},
	}
	return b
}

func StructDemo() {
	// 零值初始化
	var book Book
	fmt.Println(book, book.Author)

	// 使用复合字面量来初始化
	person := Persion{
		Name:  "peter.liu",
		Phone: "123",
		Addr:  "sh",
	}
	book2 := Book{
		Title:  "js is great, go is better",
		Author: person,
	}
	fmt.Println(book2, book2.Author)

	// 使用函数
	book3 := newBook("peter.liu", "456", "leshan", "go is good")
	fmt.Println(book3)

	// 内存长度
	var s1 S1
	var s2 S2
	fmt.Println(unsafe.Sizeof(s1)) // 24
	fmt.Println(unsafe.Sizeof(s2)) // 16
}
