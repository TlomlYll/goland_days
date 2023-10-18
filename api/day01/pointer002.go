package day01

import "fmt"

//指针传递参数

// int类型的指针作为函数的形参
func pointerVar(p *int) {
	*p += 100
}
func exchangeVar(i, j *int) {
	*i, *j = *j, *i

}
func pointer02() {
	var i int = 1
	fmt.Println(i)

	pointerVar(&i)
	fmt.Println(i)
	var x, y = 66, 99
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	//	交换之后的值
	exchangeVar(&x, &y)
	fmt.Println("交换后的x:", x)
	fmt.Println("交换后的y:", y)

}
