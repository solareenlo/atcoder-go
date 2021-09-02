package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	var s string
	var d int
	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&s, &d)
		m := 0
		if d >= a && d <= b {
			m = d
		} else if d < a {
			m = a
		} else {
			m = b
		}
		if s == "West" {
			m *= -1
		}
		res += m
	}

	if res == 0 {
		fmt.Println(0)
	} else if res > 0 {
		fmt.Println("East", res)
	} else {
		fmt.Println("West", -res)
	}
}
