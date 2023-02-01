package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [22][22]int
var n int
var M []map[int]int
var ans int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	M = make([]map[int]int, 22)
	for i := range M {
		M[i] = make(map[int]int)
	}

	A(0, 0, 0)
	B(n-1, n-1, 0)

	fmt.Println(ans)
}

func A(x, y, v int) {
	v ^= a[x][y]
	if x+y == n-1 {
		M[x][v]++
		return
	}
	A(x+1, y, v)
	A(x, y+1, v)
}

func B(x, y, v int) {
	if x+y == n-1 {
		ans += M[x][v]
		return
	}
	v ^= a[x][y]
	B(x-1, y, v)
	B(x, y-1, v)
}
