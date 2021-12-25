// Echo2 prints its command-line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", "" //这种形式只能用于函数内部，而不能用在包变量

	//range产生一对值；索引以及在该索引处的元素值
	// Go不允许使用无用的局部变量，因此可以使用 “空标识符（blank identifier）”: '_'
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
