package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x3, y3, x4, y4 int
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4)

	ABx := x2 - x1
	ABy := y2 - y1
	ACx := x3 - x1
	ACy := y3 - y1
	ADx := x4 - x1
	ADy := y4 - y1

	CAx := x1 - x3
	CAy := y1 - y3
	CBx := x2 - x3
	CBy := y2 - y3
	CDx := x4 - x3
	CDy := y4 - y3

	a := ABx*ACy - ACx*ABy
	b := ABx*ADy - ADx*ABy
	c := CDx*CAy - CAx*CDy
	d := CDx*CBy - CBx*CDy

	if a == 0 && b == 0 && c == 0 && d == 0 {
		v1 := pair{x1, y1}
		v2 := pair{x2, y2}
		v3 := pair{x3, y3}
		v4 := pair{x4, y4}
		if x1 > x2 || (x1 == x2 && y1 > y2) {
			v1, v2 = v2, v1
		}
		if x3 > x4 || (x3 == x4 && y3 > y4) {
			v3, v4 = v4, v3
		}
		tmp1 := maxPair(v1, v3)
		tmp2 := minPair(v2, v4)
		if tmp1.x < tmp2.x || (tmp1.x == tmp2.x && tmp1.y < tmp2.y) || (tmp1.x == tmp2.x && tmp1.y == tmp2.y) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}

	ok1 := (a >= 0 && b <= 0) || (a <= 0 && b >= 0)
	ok2 := (c >= 0 && d <= 0) || (c <= 0 && d >= 0)
	if ok1 && ok2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

type pair struct{ x, y int }

func maxPair(a, b pair) pair {
	if a.x > b.x || (a.x == b.x && a.y > b.y) {
		return a
	}
	return b
}

func minPair(a, b pair) pair {
	if a.x < b.x || (a.x == b.x && a.y < b.y) {
		return a
	}
	return b
}
