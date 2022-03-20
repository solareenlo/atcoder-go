package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	m := Factorization(n)

	for i := range m {
		fmt.Fprint(out, m[i], " ")
	}
}

func Factorization(n int) []int {
	num := n
	var res []int
	for i := 2; i*i <= n; i++ {
		if num%i == 0 {
			num = num / i
			res = append(res, i)
			i = 1
		}
	}

	if num > 1 {
		res = append(res, num)
	}
	return res
}
