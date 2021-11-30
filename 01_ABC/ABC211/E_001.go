package main

import "fmt"

type pair struct{ x, y int }

var (
	red    int
	used   = map[int]bool{}
	s      = make([]string, 0)
	answer int
	n      int
	dx     = [4]int{1, 0, -1, 0}
	dy     = [4]int{0, 1, 0, -1}
)

func valid(x, y int) bool {
	return 0 <= x && x <= n-1 && 0 <= y && y <= n-1
}

func change(P pair) {
	num := P.x*n + P.y
	if s[P.x][P.y] == '@' {
		red -= 1 << num
		s[P.x] = replaceAtIndex(s[P.x], '.', P.y)
	} else {
		red += 1 << num
		s[P.x] = replaceAtIndex(s[P.x], '@', P.y)
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func dfs(num int) {
	if _, ok := used[red]; ok {
		return
	}
	used[red] = true

	if num == 0 {
		answer++
		return
	}

	next := make([]pair, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == '.' {
				flag := false
				for z := 0; z <= 3; z++ {
					nxt_i := i + dx[z]
					nxt_j := j + dy[z]
					if valid(nxt_i, nxt_j) && s[nxt_i][nxt_j] == '@' {
						flag = true
					}
				}
				if flag {
					next = append(next, pair{i, j})
				}
			}
		}
	}
	for _, pos := range next {
		change(pos)
		dfs(num - 1)
		change(pos)
	}
}

func main() {
	var k int
	fmt.Scan(&n, &k)

	s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == '.' {
				change(pair{i, j})
				dfs(k - 1)
				change(pair{i, j})
			}
		}
	}

	fmt.Println(answer)
}
