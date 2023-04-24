package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, tx, ty int
	fmt.Fscan(in, &n, &tx, &ty)
	res := 0
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		res += abs(tx-x) + abs(ty-y)
		tx = x
		ty = y
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
