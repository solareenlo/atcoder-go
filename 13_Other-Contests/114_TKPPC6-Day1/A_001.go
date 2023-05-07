package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	ans := -1
	if 2015 <= N && N <= 2016 {
		ans = N - 2015 + 1
	} else if N >= 2018 {
		ans = N - 2018 + 3
	}
	fmt.Println(ans)
}
