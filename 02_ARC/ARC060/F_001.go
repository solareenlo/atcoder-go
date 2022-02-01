package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N = 500005

var (
	s string
	a = make([][]int, 2)
	n int
	b = [2][N]int{}
)

func kmp(nxt []int) {
	for i, j := 2, 0; i <= n; i++ {
		for j > 0 && s[i] != s[j+1] {
			j = nxt[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		nxt[i] = j
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &s)
	n = len(s)

	if strings.Count(s, string(s[0])) == n {
		fmt.Println(n)
		fmt.Println(1)
		return
	}

	for i := range a {
		a[i] = make([]int, N)
	}
	s = "#" + s
	kmp(a[0])
	s = "#" + reverseString(s[1:])
	kmp(a[1])
	if a[0][n] == 0 || n%(n-a[0][n]) != 0 {
		fmt.Println(1)
		fmt.Println(1)
		return
	}

	for i := 0; i < 2; i++ {
		for j := 1; j <= n; j++ {
			if a[i][j] == 0 {
				b[i][j] = 1
			} else {
				b[i][j] = j % (j - a[i][j])
			}
		}
	}

	cnt := 0
	for i := 1; i < n; i++ {
		if b[0][i] != 0 && b[1][n-i] != 0 {
			cnt++
		}
	}
	fmt.Println(2)
	fmt.Println(cnt)
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
