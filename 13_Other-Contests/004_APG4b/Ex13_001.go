package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
		sum += v[i]
	}

	ave := sum / len(v)
	for i := 0; i < n; i++ {
		if v[i] < ave {
			fmt.Println(ave - v[i])
		} else {
			fmt.Println(v[i] - ave)
		}
	}
}
