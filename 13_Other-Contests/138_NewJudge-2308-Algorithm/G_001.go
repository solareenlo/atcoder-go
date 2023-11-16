package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	b, c int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	const N = 1000010
	ele := make([][]pair, N)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		ele[a] = append(ele[a], pair{b, c})
	}

	vec := make([]pair, 0)
	for i := 1; i <= n; i++ {
		sort.Slice(ele[i], func(a, b int) bool {
			if ele[i][a].b == ele[i][b].b {
				return ele[i][a].c < ele[i][b].c
			}
			return ele[i][a].b < ele[i][b].b
		})
		for _, p := range ele[i] {
			if len(vec) != 0 && p.b <= vec[len(vec)-1].c {
				vec[len(vec)-1].c = max(vec[len(vec)-1].c, p.c)
			} else {
				vec = append(vec, p)
			}
		}
		ele[i] = ele[i][:0]
		ele[i] = append(ele[i], vec...)
		vec = vec[:0]
	}

	all := make([]pair, 0)
	for i := 1; i <= n; i++ {
		for _, p := range ele[i] {
			all = append(all, p)
		}
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].b == all[j].b {
			return all[i].c < all[j].c
		}
		return all[i].b < all[j].b
	})

	for _, p := range all {
		if len(vec) != 0 {
			if vec[len(vec)-1].c >= p.c {
				continue
			}
			if vec[len(vec)-1].b == p.b {
				vec = vec[:len(vec)-1]
			}
		}
		vec = append(vec, p)
	}
	all = all[:0]
	all = append(all, vec...)
	copy(all, vec)
	vec = vec[:0]
	n = len(all)

	var dp [N][20]int
	for i, j := 0, 0; i < n; i++ {
		for j < n && all[i].c >= all[j].b {
			j++
		}
		dp[i][0] = j - 1
	}
	for j := 1; j < 20; j++ {
		for i := 0; i < n; i++ {
			dp[i][j] = dp[dp[i][j-1]][j-1]
		}
	}

	for q > 0 {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		if b > d {
			a, c = c, a
			b, d = d, b
		}
		ans := d - b
		var get func([]pair, int) int
		get = func(vec []pair, pos int) int {
			it := lowerBound(vec, pair{pos + 1, -1})
			if it != 0 && pos <= vec[it-1].c {
				it--
				return it
			}
			return -1
		}

		it := get(ele[a], b)
		if it != -1 {
			b = max(b, ele[a][it].c)
		}

		it = get(ele[c], d)
		if it != -1 {
			d = min(d, ele[c][it].b)
		}

		if b < d {
			it = get(all, b)
			if it != -1 {
				ans += 2
				for i := 19; i >= 0; i-- {
					if d > all[dp[it][i]].c {
						ans += (1 << i)
						it = dp[it][i]
					}
				}
				if d > all[it].c {
					ans++
					it = dp[it][0]
				}
				if d > all[it].c {
					ans = -1
				}
			} else {
				ans = -1
			}
		} else {
			if a != c {
				ans++
			}
		}
		fmt.Fprintln(out, ans)
		q--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].b >= x.b
	})
	return idx
}
