package main

import "fmt"

func main() {
	var x, k int
	fmt.Scan(&x, &k)
	sum := pow(10, k)
	if x/sum == 0 {
		fmt.Print(sum)
	} else {
		fmt.Print(sum * (x/sum + 1))
	}
	fmt.Println()
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
