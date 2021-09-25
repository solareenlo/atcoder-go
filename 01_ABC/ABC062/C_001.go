package main

import "fmt"

func f(w, h int) int {
	res := int(1e9 + 7)
	for i := 1; i <= w; i++ {
		a := i * h
		b := (w - i) / 2 * h
		c := (w - i + 1) / 2 * h
		res = min(res, max(a, b, c)-min(a, b, c))
		b = (w - i) * (h / 2)
		c = (w - i) * ((h + 1) / 2)
		res = min(res, max(a, b, c)-min(a, b, c))
	}
	return res
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	fmt.Println(min(f(h, w), f(w, h)))
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
