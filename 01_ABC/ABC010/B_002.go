package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	for i, a := 0, 0; i < n; i++ {
		fmt.Scan(&a)
		mod6 := a % 6
		switch mod6 {
		case 2, 4:
			res++
		case 5:
			res += 2
		case 0:
			res += 3
		}
	}
	fmt.Println(res)
}
