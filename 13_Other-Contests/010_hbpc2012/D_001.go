package main

import "fmt"

const inf = 1 << 30

var n, m, r int
var draw int
var vis [16][16][4][16][16][4]bool
var memo [16][16][4][16][16][4]int

func main() {
	fmt.Scan(&n, &m, &r)
	ans := rec(1, 1, m, 1, 1, m)
	if draw != 0 {
		fmt.Println("Infinite")
	} else {
		if ans > 0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
		fmt.Println(abs(ans))
	}
}

func rec(a1, b1, q1, a2, b2, q2 int) int {
	if a1 == 0 && b1 == 0 {
		return 0
	}
	if a1 > b1 {
		a1, b1 = b1, a1
	}
	if a2 > b2 {
		a2, b2 = b2, a2
	}
	if vis[a1][b1][q1][a2][b2][q2] {
		draw = 1
		return 0
	}
	if memo[a1][b1][q1][a2][b2][q2] != 0 {
		return memo[a1][b1][q1][a2][b2][q2]
	}
	vis[a1][b1][q1][a2][b2][q2] = true
	w := inf
	o := 0
	if a1 != 0 && a2 != 0 {
		up(&w, &o, rec(f(a1, a2), b2, q2, a1, b1, q1))
	}
	if draw != 0 {
		return 0
	}
	if a1 != 0 && b2 != 0 {
		up(&w, &o, rec(f(a1, b2), a2, q2, a1, b1, q1))
	}
	if draw != 0 {
		return 0
	}
	if b1 != 0 && a2 != 0 {
		up(&w, &o, rec(f(b1, a2), b2, q2, a1, b1, q1))
	}
	if draw != 0 {
		return 0
	}
	if b1 != 0 && b2 != 0 {
		up(&w, &o, rec(f(b1, b2), a2, q2, a1, b1, q1))
	}
	if draw != 0 {
		return 0
	}
	if a1 == 0 && b1 > 1 && q1 != 0 {
		for i := 1; i < b1; i++ {
			up(&w, &o, rec(a2, b2, q2, i, b1-i, q1-1))
			if draw != 0 {
				return 0
			}
		}
	}
	vis[a1][b1][q1][a2][b2][q2] = false
	if w < inf {
		memo[a1][b1][q1][a2][b2][q2] = w
	} else {
		memo[a1][b1][q1][a2][b2][q2] = -o
	}
	return memo[a1][b1][q1][a2][b2][q2]
}

func f(a, b int) int {
	k := a + b
	if k < n {
		return k
	} else {
		if r != 0 {
			return k - n
		}
		return 0
	}
}

func up(w, o *int, v int) {
	if v < 1 {
		*w = min(*w, -v+1)
	} else {
		*o = max(*o, v+1)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
