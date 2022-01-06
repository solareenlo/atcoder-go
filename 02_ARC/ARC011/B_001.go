package main

import (
	"fmt"
	"strings"
)

func isSkip(c byte) bool {
	ok := false
	if c == 'a' || c == 'e' || c == 'i' {
		ok = true
	}
	if c == 'o' || c == 'u' || c == 'y' {
		ok = true
	}
	if c == ',' || c == '.' {
		ok = true
	}
	return ok
}

func main() {
	var n int
	fmt.Scan(&n)

	w := make([]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		w[i] = strings.ToLower(tmp)
	}

	num := map[byte]byte{}
	num['b'] = '1'
	num['c'] = '1'
	num['d'] = '2'
	num['w'] = '2'
	num['t'] = '3'
	num['j'] = '3'
	num['f'] = '4'
	num['q'] = '4'
	num['l'] = '5'
	num['v'] = '5'
	num['s'] = '6'
	num['x'] = '6'
	num['p'] = '7'
	num['m'] = '7'
	num['h'] = '8'
	num['k'] = '8'
	num['n'] = '9'
	num['g'] = '9'
	num['z'] = '0'
	num['r'] = '0'

	ret := make([]string, 0)
	for i := 0; i < n; i++ {
		var tmp string
		for j := 0; j < len(w[i]); j++ {
			if !isSkip(w[i][j]) {
				tmp += string(num[w[i][j]])
			}
		}
		if len(tmp) != 0 {
			ret = append(ret, tmp)
		}
	}

	N := len(ret)
	for i := 0; i < N; i++ {
		fmt.Print(ret[i])
		if i != N-1 && len(ret[i]) != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
