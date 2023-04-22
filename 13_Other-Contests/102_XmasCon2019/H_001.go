package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var T, n int
var s, t []string

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &T, &n)
	ans := 0
	t = make([]string, 100010)
	for T > 0 {
		T--
		var tmp string
		fmt.Fscan(in, &tmp)
		tmp = " " + tmp
		s = strings.Split(tmp, "")
		ans += sol()
	}
	fmt.Println(ans)
}

func sol() int {
	if strings.Count(strings.Join(s[1:n+1], ""), ".") == 0 {
		return 0
	}
	if chk(s) {
		return 1
	}
	if work() {
		return 2
	}
	return 3
}

func chk(s []string) bool {
	v := make([]int, 0)
	for i := 1; i <= n; i++ {
		if s[i] == "." {
			v = append(v, i)
		}
	}
	if len(v) < 2 {
		return true
	}
	x := 0
	for i := 1; i < len(v); i++ {
		x = gcd(v[i]-v[0], x)
	}
	for (x & 1) == 0 {
		x >>= 1
	}
	return x > 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func work() bool {
	v := make([]int, 0)
	for i := 1; i <= n; i++ {
		if s[i] == "." {
			v = append(v, i)
		}
	}
	if len(v) < 3 {
		return true
	}
	return cal(v[1]-v[0], v[0]) || cal(v[2]-v[1], v[1]) || cal(v[2]-v[0], v[0])
}

func cal(p, x int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			for p%i == 0 {
				p /= i
			}
			if i > 2 {
				for j := 1; j <= n; j++ {
					t[j] = s[j]
				}
				for j := x; j <= n; j += i {
					t[j] = "#"
				}
				if chk(t) {
					return true
				}
			}
		}
	}
	if p > 2 {
		for j := 1; j <= n; j++ {
			t[j] = s[j]
		}
		for j := x; j <= n; j += p {
			t[j] = "#"
		}
		if chk(t) {
			return true
		}
	}
	return false
}
