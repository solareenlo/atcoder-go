package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MOD = 1000000007
const MAXN = 200200

var s []string
var co, no []int
var i, rev, cn int
var a [MAXN]int

func dfs(x int) int {
	y, z := 1, 0
	for {
		switch s[i] {
		case "(":
			i++
			y = y * dfs(x*y%MOD) % MOD
		case ")":
			i++
			return (z + y) % MOD
		case "+":
			z = (z + y) % MOD
			y = 1
			i++
		case "-":
			if rev == 0 {
				z = (z + y) % MOD
				y = MOD - 1
			} else {
				z = (z - y + MOD) % MOD
				y = 1
			}
			i++
		case "*":
			i++
		case "a":
			co[no[i]] = co[no[i]] * x % MOD * y % MOD
			y = y * a[no[i]] % MOD
			i++
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	co = make([]int, MAXN)

	var tmp string
	fmt.Fscan(in, &tmp)
	n := len(tmp)
	tmp = " " + tmp + " "
	s = strings.Split(tmp, "")
	var m int
	fmt.Fscan(in, &m)
	s[0] = "("
	s[n+1] = ")"
	no = make([]int, n+2)
	for j := 1; j <= n; j++ {
		if s[j] == "a" {
			cn++
		}
		no[j] = cn
	}
	for j := 1; j <= cn; j++ {
		fmt.Fscan(in, &a[j])
		co[j] = 1
	}
	i = 1
	res := dfs(1)
	reverseOrderInt(no)
	reverseOrderString(s)
	for j := 0; j < n+2; j++ {
		if s[j] == "(" {
			s[j] = ")"
		} else if s[j] == ")" {
			s[j] = "("
		}
	}
	i = 1
	rev = 1
	dfs(1)
	for m > 0 {
		m--
		var b, x int
		fmt.Fscan(in, &b, &x)
		x = (x + MOD - a[b]) % MOD
		fmt.Fprintln(out, (res+x*co[b])%MOD)
	}
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reverseOrderString(a []string) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
