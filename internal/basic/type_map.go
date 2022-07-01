package basic

import "fmt"

func MapDemo() {
	// var m map[string]int = make(map[string]int)
	m := map[string]int{"a1": 1}
	m["a2"] = 2
	v, ok := m["a3"]
	fmt.Println(m, v, ok)
}
