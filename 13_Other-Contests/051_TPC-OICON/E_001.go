package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&v[i])
	}
	res := []int{}
	for {
		bit := 0
		for i := len(v) - 1; i >= 0; i-- {
			bit |= (1 << v[i])
			if bit == 1023 {
				bit = 0
				i--
			}
		}
		val := int(1e9)
		if len(res) == 0 {
			for i := 1; i < 10; i++ {
				if (bit>>i)%2 == 0 {
					val = min(val, i)
				}
			}
			if val == 1e9 {
				val = 1
			}
		} else {
			for i := 0; i < 10; i++ {
				if (bit>>i)%2 == 0 {
					val = min(val, i)
				}
			}
		}
		res = append(res, val)
		for {
			if len(v) == 0 {
				for _, i := range res {
					fmt.Print(i)
				}
				fmt.Println()
				return
			}
			if v[0] == val {
				v = v[1:]
				if len(v) != 0 {
					v = v[1:]
				}
				break
			}
			v = v[1:]
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
