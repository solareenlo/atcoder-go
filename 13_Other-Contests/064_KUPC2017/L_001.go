package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2222
const M = 111111

type pair struct {
	x, y int
}

var n, m int
var E [M][]pair
var vis [M]bool
var color, num [M]int
var nosol int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &m, &n)
	for i := 1; i < m+1; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = " " + s
		for j := 1; j < n+1; j++ {
			if s[j] == 'x' {
				ae(i, j+m, 1)
			}
			if s[j] == 'o' {
				ae(i, j+m, 0)
			}
		}
	}
	ans, sum, ans2, b2, b2_no1, b2_no2, bno0, b2_1_2, b2_1_4 := 0, 0, 0, 0, 0, 0, 0, 0, 0
	for i := 1; i < m+n+1; i++ {
		if !vis[i] {
			num[0] = 0
			num[1] = 0
			dfs(i)
			p := (num[0] & 1) + (num[1] & 1)
			if num[1] == 0 {
				ans ^= (num[0] & 1)
				ans2 ^= (num[0] & 1)
			} else {
				if p != 1 {
					bno0 += min(num[0], num[1])
				}
				if p == 2 {
					ans ^= 1
				}
				if p == 1 {
					ans ^= 2
					b2++
					if num[0] != 1 && num[1] != 1 {
						b2_no1++
					}
					if num[0] != 2 && num[1] != 2 {
						b2_no2++
					}
					if num[0]+num[1] == 3 {
						b2_1_2++
					} else {
						b2_1_4++
					}
				}
			}
			sum += min(num[0], num[1])
		}
	}
	if nosol != 0 {
		fmt.Println((m + n) & 1)
		return
	}
	if sum == 1 {
		fmt.Println(3)
		return
	}
	if bno0 == 0 && ans == 3 && (b2&1) != 0 && b2_no2 == 0 && (b2 == 1 || b2_no1 == 0) {
		fmt.Println(2)
		return
	}
	if bno0 == 1 && ans == 2 && (b2&1) != 0 && b2_no1 == 0 && b2_1_2 < b2_1_4+3 && ans2 != 0 {
		fmt.Println(2)
		return
	}
	if bno0 == 1 && ans == 0 && (b2&1) == 0 && b2_no1 == 0 && b2_1_2 < b2_1_4+3 && ans2 != 0 {
		fmt.Println(3)
		return
	}
	if ans == 0 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}

func ae(x, y, z int) {
	E[x] = append(E[x], pair{y, z})
	E[y] = append(E[y], pair{x, z})
}

func dfs(x int) {
	vis[x] = true
	for _, i := range E[x] {
		if !vis[i.x] {
			color[i.x] = color[x] ^ i.y
			vis[i.x] = true
			dfs(i.x)
		} else if color[i.x] != (color[x] ^ i.y) {
			nosol = 1
		}
	}
	num[color[x]]++
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
