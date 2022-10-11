package main

import (
	"fmt"
	"os"
)

// 练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。

func two() {
	for i, v := range os.Args[:] {
		fmt.Printf("Args[%v]:%v\n", i, v)
	}
}
