package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		m[a]++
	}

	res := 0
	for _, v := range m {
		res += v * (v - 1) / 2
	}
	fmt.Println(res)
}
