package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		if check(a, b, d, c) != 0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
	}
}

func check(I, J, K, L int) int {
	s := [4]int{I, J, K, L}
	for cyc := 0; cyc < 17; cyc++ {
		cur := [4]int{s[cyc%4], s[(cyc+1)%4], s[(cyc+2)%4], s[(cyc+3)%4]}
		for j := 0; j < 4; j++ {
			if s[j] != cur[j] {
				if cur[j] < s[j] {
					cur, s = s, cur
				}
				break
			}
		}
	}
	ans := -1
	if s[2] <= s[1] || s[2] <= s[3] {
		if s[0] != s[2] {
			ans = 1
		} else {
			ans = 0
		}
	} else if s[0] == 0 {
		if s[2] == 0 {
			ans = 0
		} else {
			if s[1] != s[3] {
				ans = 1
			} else {
				ans = 0
			}
		}
	} else {
		if s[1] != s[3] {
			ans = 1
		} else {
			i := s[0]
			j := s[1]
			k := s[2]
			ans = F(i, j, k)
		}
	}
	return ans
}

func F(x, y, z int) int {
	return calc(x, y-x, z-y)
}

func calc(x, a, b int) int {
	if x%(a+b) >= b {
		return 1
	}
	return 0
}
