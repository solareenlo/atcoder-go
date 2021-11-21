package main

import "fmt"

func main() {
	var h, w, x, y int
	fmt.Scan(&h, &w, &x, &y)
	x--
	y--

	s := make([]string, h)
	for i := range s {
		fmt.Scan(&s[i])
	}

	res := 0
	for i := x; i < h && s[i][y] == '.'; i++ {
		res++
	}
	for i := x; i >= 0 && s[i][y] == '.'; i-- {
		res++
	}
	for i := y; i < w && s[x][i] == '.'; i++ {
		res++
	}
	for i := y; i >= 0 && s[x][i] == '.'; i-- {
		res++
	}

	fmt.Println(res - 3)
}
