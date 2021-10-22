package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	for i := 1; i < n+1; i++ {
		s := fmt.Sprint(i)
		if len(s)%2 != 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
