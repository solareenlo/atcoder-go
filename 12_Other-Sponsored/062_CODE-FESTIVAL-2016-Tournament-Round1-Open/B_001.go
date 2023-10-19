package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var k, le, n int
var s string
var ss []string

func main() {
	in := bufio.NewReader(os.Stdin)

	var dd [100005]bool

	fmt.Fscan(in, &k, &s)
	k++
	n = len(s)
	s = " " + s
	le = n / k
	if n%k != 0 {
		le++
	}
	dd[1] = true
	for i := 1; i <= n-le+1; i++ {
		if dd[i] {
			dd[i+le-1] = true
			dd[i+le] = true
		}
	}
	if n%k == 0 {
		for i := 1; i <= n-le+1; i += le {
			if dd[i] {
				ss = append(ss, s[i:i+le])
			}
		}
		sort.Strings(ss)
		fmt.Println(ss[len(ss)-1])
		return
	}
	for i := 1; i <= n-le+1; i++ {
		if dd[i] {
			ss = append(ss, s[i:i+le])
		}
	}
	sort.Strings(ss)
	l := 0
	r := len(ss) - 1
	for l <= r {
		m := (l + r) >> 1
		if check(ss[m]) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	fmt.Println(ss[r+1])
}

func check(x string) bool {
	vt := 1
	for i := 1; i <= k; i++ {
		if vt+le-1 > n {
			return true
		}
		if x >= s[vt:vt+le] {
			vt += le
		} else {
			vt += le - 1
		}
	}
	if vt > n {
		return true
	}
	return false
}
