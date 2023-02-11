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
	m[100] = 0
	m[200] = 0
	m[300] = 0
	m[400] = 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		m[a]++
	}

	fmt.Println(m[100]*m[400] + m[200]*m[300])
}
