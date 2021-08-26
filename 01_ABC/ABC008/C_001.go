package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	c := make([]int, n)
	for i := range c {
		fmt.Scan(&c[i])
	}

	res := float64(0)
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if c[i]%c[j] == 0 {
				cnt++
			}
		}
		res += float64((cnt+1)/2) / float64(cnt)
	}
	fmt.Println(res)
}
