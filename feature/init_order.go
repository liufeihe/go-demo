package feature

import "fmt"

var (
	_  = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("const c1 has been initialized.")
	}
	if c2 != "" {
		fmt.Println("const c2 has been initialized.")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("var %s has been initialzed \n", name)
	return name
}

func init() {
	fmt.Println("first init has been called.")
}

func init() {
	fmt.Println("second init has been called.")
}

func Hello() {
	fmt.Printf("hello, c1: %s, c2: %s \n", c1, c2)
}

func main() {
	fmt.Println("main.")
}
