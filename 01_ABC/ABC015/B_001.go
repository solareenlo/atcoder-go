package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a, num, sum := 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += a
		if a != 0 {
			num++
		}
	}

	if sum%num == 0 {
		fmt.Println(sum / num)
	} else {
		fmt.Println(sum/num + 1)
	}
}
