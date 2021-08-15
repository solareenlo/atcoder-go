package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 4
	switch {
	case 126 <= n && n <= 211:
		res = 6
	case 212 <= n && n <= 214:
		res = 8
	}
	fmt.Println(res)
}
