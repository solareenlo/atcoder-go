package main

import "fmt"

var sx, sy int

func pt(x, y int) {
	fmt.Println(x, y)
	sx = (x << 1) - sx
	sy = (y << 1) - sy
}

func main() {
	var tx, ty, a, b, c, d int
	fmt.Scan(&sx, &sy)
	fmt.Scan(&tx, &ty)
	fmt.Scan(&a, &b, &c, &d)
	if ((sx+tx)&1) != 0 || ((sy+ty)&1) != 0 {
		fmt.Println("No")
		return
	}
	even := (a != b || sx == tx) && (c != d || sy == ty)
	odd := (a != b || sx+tx == a+b) && (c != d || sy+ty == c+d)
	if !even && !odd {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	if !even {
		pt(a, c)
	}
	for sx < tx {
		pt(a, c)
		pt(a+1, c)
	}
	for sx > tx {
		pt(a+1, c)
		pt(a, c)
	}
	for sy < ty {
		pt(a, c)
		pt(a, c+1)
	}
	for sy > ty {
		pt(a, c+1)
		pt(a, c)
	}
}
