package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var I int
var mp map[byte]pair

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < 31; i++ {
		I |= 1 << i
	}
	var N int
	fmt.Fscan(in, &N)
	mp = make(map[byte]pair)
	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(in, &s)
		c := s[0]
		ind := 0
		ans := definition(s, &ind)
		mp[c] = ans
		fmt.Fprintf(out, "%c(x)=((%d&x)^%d)\n", c, ans.x, ans.y)
	}
}

func definition(s string, ind *int) pair {
	*ind += 5
	return expr(s, ind)
}

func expr(s string, ind *int) pair {
	if s[*ind] == 'x' {
		*ind++
		return pair{I, 0}
	}
	if isdigit(s[*ind]) {
		return number(s, ind)
	}
	if isalpha(s[*ind]) {
		c := s[*ind]
		*ind++
		x := 1
		if s[*ind] == '^' {
			*ind++
			x = number(s, ind).y
		}
		*ind++
		ret := expr(s, ind)
		*ind++
		if x == 0 {
			return ret
		}
		f := mp[c]
		for i := 0; i < (x&1)+2; i++ {
			ret = pair{f.x & ret.x, (f.x & ret.y) ^ f.y}
		}
		return ret
	}
	*ind++
	X := expr(s, ind)
	c := s[*ind]
	*ind++
	Y := expr(s, ind)
	*ind++
	return op(X, Y, c)
}

func op(a, b pair, c byte) pair {
	if c == '^' {
		return pair{a.x ^ b.x, a.y ^ b.y}
	}
	if c == '&' {
		return pair{(a.x & b.x) ^ (a.y & b.x) ^ (a.x & b.y), a.y & b.y}
	}
	ret := op(a, b, '&')
	ret = op(ret, a, '^')
	ret = op(ret, b, '^')
	return ret
}

func number(s string, ind *int) pair {
	t := 0
	for *ind != len(s) && isdigit(s[*ind]) {
		t *= 10
		t += int(s[*ind] - '0')
		*ind++
	}
	return pair{0, t}
}

func isdigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isalpha(c byte) bool {
	return 'a' <= c && c <= 'z'
}
