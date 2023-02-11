package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	cnt := 0
	for i := 1; i < n+1; i++ {
		if i%x == 0 || i%y == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
