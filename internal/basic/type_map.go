package basic

import (
	"fmt"
	"time"
)

type Position struct {
	x float64
	y float32
}

func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d, %d]", k, v)
	}
}

func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}

func concurrent() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()
	time.Sleep(5 * time.Second)
}

func MapDemo() {
	// var m map[string]int = make(map[string]int)
	var m map[string]int = make(map[string]int, 8)
	fmt.Println(m)

	m2 := map[Position]string{
		{29.1, 52.1}: "js",
		{29.2, 52.2}: "java",
		{29.3, 52.3}: "go",
	}
	fmt.Println(m2[Position{29.1, 52.1}])

	m3 := map[string]int{"a1": 1}
	m3["a2"] = 2
	v, ok := m3["a3"]
	fmt.Println(m3, v, ok)
	delete(m3, "a2")
	fmt.Println(m3)

	m3["a4"] = 4
	fmt.Printf("{")
	for k, v := range m3 { // 不要依赖遍历map得到的顺序
		fmt.Printf("[%s, %d]", k, v)
	}
	fmt.Printf("}\n")

	// concurrent()
}
