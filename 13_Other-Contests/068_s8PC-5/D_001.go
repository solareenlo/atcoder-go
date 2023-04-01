package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	tmp := 0
	if h == 1 || w == 1 {
		tmp = 1
	}
	f := tmp ^ h&w&1
	if f != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}

	p, q := 1, 1
	var a [55][55]int
	a[1][1] = 1
	x, y := 1, 1
	if f == 0 {
		fmt.Scan(&x, &y)
	}
	for x+y < h+w-1 {
		a[x][y] = 1
		for a[p][q] != 0 {
			if p < h && q > 1 {
				p++
				q--
				continue
			}
			q += p
			p = 1
			if q > w {
				p += q - w
				q = w
			}
		}
		a[p][q] = 1
		fmt.Println(p, q)
		fmt.Scan(&x, &y)
	}
	fmt.Println(h, w)
	fmt.Scan(&x, &y)
}
