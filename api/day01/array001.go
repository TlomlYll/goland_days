package day01

import "fmt"

// 数组  值 地址
func array001() {
	var a = [5]int{1, 2, 3, 4, 5}
	for i := range a {
		fmt.Println(&i)
	}
	//查看数组中的元素的地址
	L := len(a)
	for i := 0; i < L; i++ {
		fmt.Println(&a[i])
	}

}
