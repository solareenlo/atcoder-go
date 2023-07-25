package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 50000
	const B = 230

	var n, s1, s2 int
	fmt.Fscan(in, &n, &s1, &s2)

	v := make([][]int, 100010)
	w := make([][]int, 100010)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		v[a] = append(v[a], b)
		w[b] = append(w[b], a)
	}

	for i := 1; i <= MX; i++ {
		sort.Ints(v[i])
		sort.Ints(w[i])
	}

	ans := 0
	for i := 1; i <= MX; i++ {
		for _, x := range v[i] {
			for j := i + 1; j < i+B; j++ {
				le := (s1 + j - i - 1) / (j - i)
				ri := s2 / (j - i)
				num := upperBound(v[j], x+ri) - lowerBound(v[j], x+le)
				ans += max(num, 0)
			}
		}
	}
	for i := 1; i <= MX; i++ {
		for _, x := range w[i] {
			for j := i + 1; j < i+B; j++ {
				le := (s1 + j - i - 1) / (j - i)
				ri := s2 / (j - i)
				le = max(B, le)
				num := upperBound(w[j], x+ri) - lowerBound(w[j], x+le)
				ans += max(num, 0)
			}
		}
	}
	fmt.Println(ans)

}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
