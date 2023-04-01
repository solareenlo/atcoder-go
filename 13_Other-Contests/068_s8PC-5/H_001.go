package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const INF = int(1e18) + 10

var S string
var N, Q int
var link, skip, skip_cnt [312345][26]int
var cnt [312345]int
var t int
var moji [1123456]byte
var Len [1123456]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &S, &Q)
	N = len(S)
	build()
	for j := 0; j < Q; j++ {
		var K, p int
		fmt.Fscan(in, &K, &p)
		query(K)
		if t == 0 {
			fmt.Fprintln(out, -1)
			continue
		}
		l := 0
		for i := 0; i < t; i++ {
			l += Len[i]
		}
		if l < p {
			p = l
		}
		ret := make([]string, p)
		for i := range ret {
			ret[i] = "a"
		}
		l = 0
		for i := t - 1; i >= 0 && p >= 1; i-- {
			for j := 0; j < Len[i] && p >= 1; j++ {
				p--
				ret[p] = string(moji[i])
			}
		}
		fmt.Fprintln(out, strings.Join(ret, ""))
	}
}

func build() {
	for i := 0; i < 26; i++ {
		link[N][i] = -1
	}
	cnt[N] = 1
	for ri := 0; ri < N; ri++ {
		i := N - ri - 1
		for j := 0; j < 26; j++ {
			link[i][j] = link[i+1][j]
		}
		link[i][S[i]-'a'] = i + 1
		cnt[i] = 1
		for j := 0; j < 26; j++ {
			if link[i][j] != -1 {
				cnt[i] = min(cnt[i]+cnt[link[i][j]], INF)
			}
		}
		for j := 0; j < 26; j++ {
			skip[i+1][j] = -1
		}
		if link[i+1][S[i]-'a'] != -1 {
			skip[i+1][0] = link[i+1][S[i]-'a']
			skip_cnt[i+1][0] = 1
			for j := 0; j < int(S[i]-'a'); j++ {
				if link[i+1][j] != -1 {
					skip_cnt[i+1][0] = min(skip_cnt[i+1][0]+cnt[link[i+1][j]], INF)
				}
			}
			k := 0
			for skip[skip[i+1][k]][k] != -1 {
				skip[i+1][k+1] = skip[skip[i+1][k]][k]
				skip_cnt[i+1][k+1] = min(skip_cnt[i+1][k]+skip_cnt[skip[i+1][k]][k], INF)
				k++
			}
		}
	}
}

func query(K int) {
	t = 0
	for i := 0; i < 26; i++ {
		if link[0][i] != -1 {
			if K <= cnt[link[0][i]] {
				dfs(K, link[0][i])
				return
			} else {
				K -= cnt[link[0][i]]
			}
		}
	}
}

func dfs(K, id int) {
	moji[t] = S[id-1]
	if K == 1 {
		Len[t] = 1
		t++
	} else {
		if skip[id][0] != -1 && skip_cnt[id][0] < K && K-skip_cnt[id][0] <= cnt[skip[id][0]] {
			for i := 1; i < 26; i++ {
				if skip[id][i] == -1 || skip_cnt[id][i] >= K || K-skip_cnt[id][i] > cnt[skip[id][i]] {
					Len[t] = 1 << (i - 1)
					t++
					dfs(K-skip_cnt[id][i-1], skip[id][i-1])
					return
				}
			}
		} else {
			K--
			Len[t] = 1
			t++
			for i := 0; i < 26; i++ {
				if link[id][i] != -1 {
					if cnt[link[id][i]] >= K {
						dfs(K, link[id][i])
						return
					} else {
						K -= cnt[link[id][i]]
					}
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
