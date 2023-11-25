package main

import (
	"fmt"
	"strconv"
)

var s string
var cur int = 0

func number() int {
	num, _ := strconv.Atoi(string(s[cur]))
	for cur+1 < len(s) {
		if !isDigit(s[cur+1]) {
			break
		}
		cur++
		num *= 10
		tmp, _ := strconv.Atoi(string(s[cur]))
		num += tmp
	}
	cur++
	return num
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func factor() pair {
	if s[cur] == '(' {
		cur++
		fct := expr()
		cur++
		return fct
	}
	if s[cur] == 'l' {
		cur += 4
		fct := expr()
		cur++
		if fct.x > 0 {
			fct = pair{0, 2}
		} else if fct.y > 0 {
			fct.y = 1
		} else {
			fct.y = 0
		}
		fct.x = 0
		if cur >= len(s) {
			return fct
		}
		if s[cur] == '^' {
			cur++
			num := number()
			fct.x *= num
			if fct.y%2 != 0 {
				fct.y = (fct.y-1)*num + 1
			} else {
				fct.y *= num
			}
		}
		return fct
	}
	if s[cur] == 'N' {
		cur++
		if cur >= len(s) {
			return pair{1, 0}
		}
		if s[cur] == '^' {
			cur++
			num := number()
			return pair{num, 0}
		} else {
			return pair{1, 0}
		}
	}
	fmt.Println(s[cur])
	return pair{0, 0}
}

func term() pair {
	tm := factor()
	if cur >= len(s) {
		return tm
	}
	if s[cur] == '*' {
		cur++
		tm2 := term()
		tm.x += tm2.x
		if tm.y%2 == 1 && tm2.y%2 == 1 {
			tm.y += tm2.y - 1
		} else {
			tm.y += tm2.y
		}
	}
	return tm
}

func expr() pair {
	ex := term()
	if cur >= len(s) {
		return ex
	}
	if s[cur] == '+' {
		cur++
		ex2 := expr()
		ex = maxPair(ex, ex2)
	}
	return ex
}

func main() {
	fmt.Scan(&s)
	ans := expr()
	fmt.Println(ans.x, (ans.y+1)/2)
}

type pair struct {
	x, y int
}

func maxPair(a, b pair) pair {
	if a.x == b.x {
		if a.y > b.y {
			return a
		}
		return b
	}
	if a.x > b.x {
		return a
	}
	return b
}
