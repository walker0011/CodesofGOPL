//go:build ignore
// +build ignore

// Dup1 prints the text of each line that appears more than once
// in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	// bufio.Scanner类型的变量：读取输入并将其拆成行或者单词
	// 通常是处理行形式的输入最简单的方法。
	inputs := bufio.NewScanner(os.Stdin)

	for inputs.Scan() {
		counts[inputs.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s \n", n, line)
		}
	}

}
