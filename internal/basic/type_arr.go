package basic

import "fmt"

func ArrDemo() {
	// 数组
	var arr1 [5]int
	var arr2 = [6]int{1, 2, 3, 4, 5, 6}
	var arr3 = [...]int{1, 2, 3}
	var arr4 = [...]int{9: 9}
	fmt.Println(arr1, arr2, arr3, arr4, len(arr4))
	for _, v := range arr4 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("%d ", arr4[i])
	}

	// 切片
	var slice1 = []int{1, 2, 3}
	fmt.Println(len(slice1))
	slice1 = append(slice1, 4)
	fmt.Println(len(slice1))
	// 使用make来创建切片
	sl := make([]byte, 6, 10)
	fmt.Println(sl)
	// 基于数组创建切片
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 := arr[3:7:9]
	fmt.Println(arr, sl2)
	// 基于切片创建切片
	sl3 := sl2[1:2:3]
	fmt.Println(sl3)
}
