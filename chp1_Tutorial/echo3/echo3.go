package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 当连接的数据量巨大时，这种方式简单高效
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
