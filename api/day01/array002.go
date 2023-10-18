package day01

import "fmt"

// const 定义常量，不可修改，类比java中的final
const N = 10

func array002() {

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
