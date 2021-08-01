package main

import "fmt"

func main() {
	x := make([]int, 4)
	fmt.Scanf("%1d%1d%1d%1d", &x[0], &x[1], &x[2], &x[3])

	res := "Strong"
	if x[0] == x[1] && x[0] == x[2] && x[0] == x[3] {
		res = "Weak"
	}
	if (x[0]+1)%10 == x[1] && (x[0]+2)%10 == x[2] && (x[0]+3)%10 == x[3] {
		res = "Weak"
	}
	fmt.Println(res)
}
