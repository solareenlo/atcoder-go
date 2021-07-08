package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	sum := 0
	for sum < n {
		sum += cnt
		cnt++
	}
	fmt.Println(cnt - 1)
}
