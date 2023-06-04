package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [1 << 15]int
	var ok [1 << 15]bool

	var N int
	fmt.Fscan(in, &N)
	A := make([]pair, 0)
	for i := 0; i < N; i++ {
		var c string
		fmt.Fscan(in, &c)
		tmp := strings.Split(c, ":")
		h1, _ := strconv.Atoi(tmp[0])
		m1, _ := strconv.Atoi(tmp[1])
		fmt.Fscan(in, &c)
		tmp = strings.Split(c, ":")
		h2, _ := strconv.Atoi(tmp[0])
		m2, _ := strconv.Atoi(tmp[1])
		A = append(A, pair{h1*60 + m1, h2*60 + m2})
	}
	sortPair(A)
	for i := 1; i < 1<<N; i++ {
		dp[i] = 114514
		id := -1
		flag := true
		p := 0
		var f int
		for j := 0; j < N; j++ {
			if ((i >> j) & 1) != 0 {
				if id < 0 {
					id = j
					f = A[j].x
				}
				if p <= A[j].x {
					p = A[j].y
				} else {
					flag = false
					break
				}
			}
		}
		ok[i] = flag && (p < 1440 || p%1440 <= f)
	}
	for i := 0; i < (1 << N); i++ {
		ne := ((1 << N) - 1) & ^i
		for j := ne; j != 0; j = (j - 1) & ne {
			if ok[j] {
				dp[i|j] = min(dp[i|j], dp[i]+1)
			}
		}
	}
	fmt.Println(dp[(1<<N)-1])
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
