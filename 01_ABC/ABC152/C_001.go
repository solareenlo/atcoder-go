package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mini := 300000
	cnt := 0
	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
		if p < mini {
			mini = p
			cnt++
		}
	}

	fmt.Println(cnt)
}
