package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	maxi := 0
	res := ""
	mp := map[string]int{}
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		mp[s]++
		if maxi < mp[s] {
			maxi = mp[s]
			res = s
		}
	}
	fmt.Println(res)
}
