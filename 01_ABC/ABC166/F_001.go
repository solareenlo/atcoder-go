package main

import (
	"fmt"
	"os"
)

func main() {
	var n, a, b, c int
	fmt.Scan(&n, &a, &b, &c)

	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	res := make([]byte, 0)
	rem := []int{a, b, c}

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			fmt.Println("Yes")
			for i := range res {
				fmt.Println(string(res[i]))
			}
			os.Exit(0)
		}
		x := s[idx][0] - 'A'
		y := s[idx][1] - 'A'
		for i := 0; i < 2; i++ {
			if rem[x] > 0 {
				res = append(res, 'A'+y)
				rem[x]--
				rem[y]++
				dfs(idx + 1)
				rem[x]++
				rem[y]--
				res = res[:len(res)-1]
			}
			x, y = y, x
		}
	}

	dfs(0)
	fmt.Println("No")
}
