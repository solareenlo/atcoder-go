package main

import "fmt"

type pair struct{ x, y int }

var k int

func cal(x, y int) pair {
	if x > y {
		t := cal(y, x)
		return pair{t.y, t.x}
	}
	if x < 0 {
		t := cal(-x, y)
		return pair{-t.x, t.y}
	}
	if x+y >= 2*k {
		return pair{x, y - k}
	}
	if x+y == k {
		return pair{0, 0}
	}
	return pair{(x+y)/2 - k, (x + y + 1) / 2}
}

func main() {
	var x, y int
	fmt.Scan(&k, &x, &y)

	if k&1 == 0 && (x+y)&1 != 0 {
		fmt.Println(-1)
		return
	}

	a := [202020]pair{}
	a[0].x = x
	a[0].y = y

	i := 0
	for ; a[i].x != 0 || a[i].y != 0; i++ {
		a[i+1] = cal(a[i].x, a[i].y)
	}

	fmt.Println(i)
	i--
	for ; i >= 0; i-- {
		fmt.Println(a[i].x, a[i].y)
	}
}
