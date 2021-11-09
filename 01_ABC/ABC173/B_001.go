package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := map[string]int{}
	res["AC"] = 0
	res["WA"] = 0
	res["TLE"] = 0
	res["RE"] = 0

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		res[tmp]++
	}

	fmt.Println("AC x", res["AC"])
	fmt.Println("WA x", res["WA"])
	fmt.Println("TLE x", res["TLE"])
	fmt.Println("RE x", res["RE"])
}
