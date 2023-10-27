package main

import "fmt"

func main() {
	var ans [1005]int
	ans[1000] = 1789997546303

	var p int
	fmt.Scan(&p)
	for i := 1000 - 1; i >= p; i-- {
		if ans[i+1]%2 == 0 {
			ans[i] = ans[i+1] / 2
		} else {
			ans[i] = ans[i+1]*3 + 1
		}
	}
	fmt.Println(ans[p])
}
