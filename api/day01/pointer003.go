package day01

import (
	"flag"
	"fmt"
)

// 使用指针变量获取命令行输入信息
var mode = flag.String("mode", "", "process mode")

func pointer003() {
	//	解析命令行参数，将结果写入创建的指令变量中
	flag.Parse()
	if *mode == "fast" {
		fmt.Println("fast mode execute")
	} else if *mode == "slow" {
		fmt.Println("slow mode execute")
	} else {
		fmt.Println("default mode execute")
	}
}
