// 使用指针变量获取命令行输入信息
package main

import (
	"fmt"
)

const N = 10

func main() {

	var intArr [3]int
	var stringArr [3]string
	var float64Arr [3]float64
	var float32Arr [3]float32
	var ageArr [N * 5]int
	fmt.Println(intArr)
	fmt.Println(stringArr)
	fmt.Println(float64Arr)
	fmt.Println(float32Arr)
	fmt.Println(ageArr)

}
