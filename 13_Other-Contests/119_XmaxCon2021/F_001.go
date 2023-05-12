package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	nw := "ACCAABACCCBAAACACCBBABABBCACCCCABAAAABCBACCBBCACBABAACCACBCBBAABACBBBCBBCBB"
	var ans, meow, a, b string
	for i := 0; i < len(nw); i++ {
		if nw[i] == 'A' {
			ans += "EAEE"
			meow += "AE"
			a += "RLL"
			b += "R"
		}
		if nw[i] == 'B' {
			ans += "BE"
			meow += "EBEE"
			a += "L"
			b += "RLL"
		}
		if nw[i] == 'C' {
			ans += "CE"
			meow += "CE"
			a += "R"
			b += "L"
		}
	}
	ans = ans[:len(ans)-1]
	meow = meow[:len(meow)-1]
	tmp := reverseOrderString(strings.Split(meow, ""))
	fmt.Fprintln(out, "D"+ans+"D"+strings.Join(tmp, "")+"D")
}

func reverseOrderString(a []string) []string {
	n := len(a)
	res := make([]string, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
