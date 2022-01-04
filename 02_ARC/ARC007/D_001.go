package main

import (
	"fmt"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(s, t string) string {
	ls := len(s)
	lt := len(t)
	val := make([]int, max(ls, lt)+1)
	for i := 0; i < ls; i++ {
		val[ls-i-1] += int(s[i] - '0')
	}
	for i := 0; i < lt; i++ {
		val[lt-i-1] += int(t[i] - '0')
	}
	for i := 0; i < len(val); i++ {
		if val[i] >= 10 {
			val[i] -= 10
			val[i+1] += 1
		}
	}
	for val[len(val)-1] == 0 {
		val = val[:len(val)-1]
	}
	var ans string
	for i := len(val) - 1; i >= 0; i-- {
		ans += string(val[i] + '0')
	}
	return ans
}

func subtract(s, t string) string {
	ls := len(s)
	lt := len(t)
	val := make([]int, ls)
	for i := 0; i < ls; i++ {
		val[ls-i-1] += int(s[i] - '0')
	}
	for i := 0; i < lt; i++ {
		val[lt-i-1] -= int(t[i] - '0')
	}
	for i := 0; i < len(val); i++ {
		if val[i] < 0 {
			val[i] += 10
			val[i+1] -= 1
		}
	}
	for val[len(val)-1] == 0 {
		val = val[:len(val)-1]
	}
	var ans string
	for i := len(val) - 1; i >= 0; i-- {
		ans += string(val[i] + '0')
	}
	return ans
}

func main() {
	var s string
	fmt.Scan(&s)
	if s[0] == '0' {
		s = "1" + s
	}
	sp := len(s)
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			sp = i
			break
		}
	}
	ini := s[0:sp]
	s = s[sp:]
	n := len(s)
	var ans string
	for i := sp; i < n; i++ {
		nxt := s[0:i]
		if sp < i || ini < nxt {
			diff := subtract(nxt, ini)
			cur := nxt
			var str string
			for len(str) < n {
				str += cur
				cur = add(cur, diff)
			}
			if str[0:n] == s {
				ans = diff
				break
			}
		}
	}
	if n == 0 {
		fmt.Println(ini, 1)
	} else if len(ans) == 0 {
		if sp < n || (sp == n && ini < s) {
			fmt.Println(ini, subtract(s, ini))
		} else {
			ls := s + strings.Repeat("0", sp-n)
			rs := s + strings.Repeat("9", sp-n)
			if ini >= rs {
				fmt.Println(ini, subtract(ls+"0", ini))
			} else if ini >= ls {
				fmt.Println(ini, 1)
			} else {
				fmt.Println(ini, subtract(ls, ini))
			}
		}
	} else {
		fmt.Println(ini, ans)
	}
}
