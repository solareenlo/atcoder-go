package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	s += " "

	n, cnt, sum := 0, 0, 0
	for ; s[n] != ' '; n++ {
		if n != 0 && s[n-1] != s[n] {
			cnt++
		}
		sum += int(s[n] - 'a')
		sum %= 3
	}

	mod := 998244353
	ans := 0
	if cnt == 0 {
		ans = 1
	} else {
		pw := 1
		ans = 1
		for i := 1; i < n; i++ {
			pw = pw * 2 % mod
			ans = ans * 3 % mod
		}
		tmp := 0
		if cnt == n-1 {
			tmp = 1
		}
		ans += mod - pw + tmp
		ans %= mod
		if n%3 == 0 {
			pw = 1
			for i := 1; i < n/3; i++ {
				pw = pw * 2 % mod
			}
			if sum == 0 {
				tmp = 0
				if n == 3 {
					tmp = mod - 1
				}
				pw = ((mod-pw)*2 + tmp) % mod
			}
			ans = (ans + pw) % mod
		}
	}
	fmt.Println(ans)
}
