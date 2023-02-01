package main

import "fmt"

func main() {
	x, y, c, d := 0, 0, 0, 0
	for p := 1; p <= 10; p++ {
		var ch string
		fmt.Scan(&ch)
		for q := 1; q <= 10; q++ {
			if x == 0 && c == 0 && ch[q-1] == '#' {
				x = p
				c = q
			}
			if ch[q-1] == '#' {
				y = p
				d = q
			}
		}
	}
	fmt.Println(x, y)
	fmt.Println(c, d)
}
