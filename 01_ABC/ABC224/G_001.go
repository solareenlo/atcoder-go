package main

import "fmt"

func main() {
	var n, s, t, a, b float64
	fmt.Scan(&n, &s, &t, &a, &b)

	i := 0.0
	x := (t - s) * a
	for i < t && i < 5e6 {
		s = i*a/2 + n*b/(i+1.0)
		if x < 0 || x > s {
			x = s
		}
		i += 1.0
	}
	fmt.Println(x)
}
