package day01

import "fmt"

func Pointer() {

	var n int = 99
	fmt.Println("int类型变量n的地址：", &n)
	//int指针
	var p *int
	fmt.Println("int指针p的地址：", &p)
	//int指针变量p初始化赋值，将n的值赋值给p
	p = &n
	fmt.Println("指针变量p的值是：", p)

	fmt.Println("____________________________通过指针修改n的值")
	*p = 100
	fmt.Println("指针重新赋值之后，n的值：", n)

}
