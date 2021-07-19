package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}

	candy := make(map[int]int)
	for i := 0; i < k; i++ {
		candy[c[i]]++
	}

	res := len(candy)
	for i := k; i < n; i++ {
		candy[c[i-k]]--
		if candy[c[i-k]] == 0 {
			delete(candy, c[i-k])
		}
		candy[c[i]]++
		if res < len(candy) {
			res = len(candy)
		}
	}
	fmt.Println(res)
}
