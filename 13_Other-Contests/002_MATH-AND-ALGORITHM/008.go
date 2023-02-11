package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	cnt := 0
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i+j <= s {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
