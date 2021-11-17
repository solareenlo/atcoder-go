package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n *= 2

	cnt := 0
	for i := 1; i*i < n+1; i++ {
		if n%i == 0 && (i%2 != 0 || (n/i)%2 != 0) {
			cnt++
		}
	}

	fmt.Println(cnt * 2)
}
