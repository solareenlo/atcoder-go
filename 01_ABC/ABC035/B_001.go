package main

import "fmt"

func main() {
	var s string
	var t int
	fmt.Scan(&s, &t)

	x, y, cnt := 0, 0, 0
	for i := range s {
		switch s[i] {
		case '?':
			cnt++
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	res := x + y
	if t == 1 {
		fmt.Println(res + cnt)
	} else {
		if cnt <= res {
			fmt.Println(res - cnt)
		} else {
			fmt.Println((cnt - res) % 2)
		}
	}
}
