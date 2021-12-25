// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// go中只有i ++没有++i，且i++是语句而不是表达式，因此 j = i ++非法
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
