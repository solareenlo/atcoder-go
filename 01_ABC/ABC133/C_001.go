package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	divL, remL := l/2019, l%2019
	divR, remR := r/2019, r%2019

	if divL == divR {
		mini := 2019
		for i := remL; i < remR; i++ {
			for j := remL + 1; j < remR+1; j++ {
				mini = min(mini, (i*j)%2019)
			}
		}
		fmt.Println(mini)
		return
	}
	fmt.Println(0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
