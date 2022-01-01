package main

import (
	"fmt"
	"strings"
)

var (
	s   = [8]string{}
	mx  = [8]bool{}
	my  = [8]bool{}
	mxy = [16]bool{}
	myx = [16]bool{}
)

func dfs(d int) bool {
	if d == 5 {
		return true
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if mx[j] || my[i] || mxy[i+j] || myx[i-j+7] {
				continue
			}
			mx[j], my[i], mxy[i+j], myx[i-j+7] = true, true, true, true
			tmp := strings.Split(s[i], "")
			tmp[j] = "Q"
			s[i] = strings.Join(tmp, "")
			if dfs(d + 1) {
				return true
			}
			mx[j], my[i], mxy[i+j], myx[i-j+7] = false, false, false, false
			tmp = strings.Split(s[i], "")
			tmp[j] = "."
			s[i] = strings.Join(tmp, "")
		}
	}
	return false
}

func main() {
	for i := 0; i < 8; i++ {
		fmt.Scan(&s[i])
	}
	ans := true
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if s[i][j] == 'Q' {
				if mx[j] || my[i] || mxy[i+j] || myx[i-j+7] {
					ans = false
				}
				mx[j], my[i], mxy[i+j], myx[i-j+7] = true, true, true, true
			}
		}
	}
	if ans && dfs(0) {
		for i := 0; i < 8; i++ {
			fmt.Println(s[i])
		}
	} else {
		fmt.Println("No Answer")
	}
}
