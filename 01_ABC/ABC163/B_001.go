package main

import "fmt"

func main() {
	var days, num int
	fmt.Scan(&days, &num)

	a := make([]int, num)
	for i := 0; i < num; i++ {
		var l int
		fmt.Scan(&l)
		a[i] = l
		days -= a[i]
	}

	if days < -1 {
		days = -1
	}

	fmt.Println(days)
}
