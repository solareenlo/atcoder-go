package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	ans := make([]int, 0)
	p := 1
	for N > 0 {
		if N%(p*3) == p {
			ans = append(ans, p)
			N -= p
		} else if N%(p*3) == p*2 {
			ans = append(ans, -p)
			N += p
		}
		p *= 3
	}

	fmt.Println(len(ans))
	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i], " ")
	}
}
