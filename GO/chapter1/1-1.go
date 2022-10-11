package main

import (
	"fmt"
	"os"
)

// 练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字。
func one() {
	fmt.Printf("os.Args[0]: %v\n", os.Args[0])
}
