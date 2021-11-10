package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n    int
	val  = [52]int{}
	vis  = [52]int{}
	lenA int
	lenB int
	LEN  = [52]int{}
	s    = [52]string{}
	t    = [52][52]byte{}
	a    = [1000002]byte{}
	b    = [1000002]byte{}
	c    = [1000002]byte{}
	res  int
	sum  int
)

func dfs() {
	if sum > res {
		return
	}
	if lenA != 0 && lenB != 0 && abs(lenA-lenB) <= 1 {
		res = sum
		return
	}
	ok := true
	tt := 0
	for i := 1; i <= lenA; i++ {
		tt++
		c[tt] = a[i]
	}
	for i := lenB; i > 0; i-- {
		tt++
		c[tt] = b[i]
	}
	for i := 1; i <= tt/2; i++ {
		if c[i] != c[tt-i+1] {
			ok = false
			break
		}
	}
	if tt != 0 && ok {
		res = sum
		return
	}
	if lenA <= lenB {
		for i := 1; i <= n; i++ {
			ok := true
			for j := lenA + 1; j <= min(lenB, lenA+LEN[i]); j++ {
				if b[j] != s[i][j-lenA] {
					ok = false
					break
				}
			}
			if ok {
				if vis[i] > 100 {
					continue
				}
				vis[i]++
				tmp := lenA
				sum += val[i]
				for j := lenA + 1; j <= lenA+LEN[i]; j++ {
					a[j] = s[i][j-lenA]
				}
				lenA += LEN[i]
				dfs()
				lenA = tmp
				vis[i]--
				sum -= val[i]
			}
		}
	} else {
		for i := 1; i <= n; i++ {
			ok := true
			for j := lenB + 1; j <= min(lenA, lenB+LEN[i]); j++ {
				if a[j] != t[i][j-lenB] {
					ok = false
					break
				}
			}
			if ok {
				if vis[i] > 100 {
					continue
				}
				vis[i]++
				tmp := lenB
				sum += val[i]
				for j := lenB + 1; j <= lenB+LEN[i]; j++ {
					b[j] = t[i][j-lenB]
				}
				lenB += LEN[i]
				dfs()
				lenB = tmp
				vis[i]--
				sum -= val[i]
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	res = 1 << 60
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		LEN[i] = len(s[i])
		s[i] = " " + s[i] + " "
		ok := true
		for j := 1; j <= LEN[i]; j++ {
			t[i][j] = s[i][LEN[i]-j+1]
			if t[i][j] != s[i][j] {
				ok = false
			}
		}
		fmt.Fscan(in, &val[i])
		if ok {
			res = min(res, val[i])
		}
	}
	dfs()
	if res < 1<<60 {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
