package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 303

var (
	n   int
	u   = make([]int, N)
	d   = make([]int, N)
	l   = make([]int, N)
	r   = make([]int, N)
	dir = [N][N]int{}
)

func match() {
	type pair struct{ x, y int }
	remain := make([]pair, 0)
	for i := 0; i < n; i++ {
		s := u[i] + d[i]
		if s > n {
			fmt.Println("NO")
			os.Exit(0)
		}
		remain = append(remain, pair{-(n - s), i})
	}
	sort.Slice(remain, func(i, j int) bool {
		return remain[i].x < remain[j].x
	})
	for i := 0; i < n; i++ {
		s := l[i] + r[i]
		if s > n {
			fmt.Println("NO")
			os.Exit(0)
		}
		for j := 0; j < s; j++ {
			remain[j].x++
			if remain[j].x > 0 {
				fmt.Println("-1")
				os.Exit(0)
			}
			dir[i][remain[j].y] = 1
		}
		sort.Slice(remain, func(i, j int) bool {
			return remain[i].x < remain[j].x
		})
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dir[i][j] != 0 {
				l[i]--
				if l[i] >= 0 {
					dir[i][j] = 0
				} else {
					dir[i][j] = 2
				}
			} else {
				u[j]--
				if u[j] >= 0 {
					dir[i][j] = 1
				} else {
					dir[i][j] = 3
				}
			}
		}
	}
}

var (
	used    = [N][N]bool{}
	instack = [N][N]bool{}
	dx      = [4]int{0, -1, 0, 1}
	dy      = [4]int{-1, 0, 1, 0}
)

func dfs(i, j int) bool {
	if used[i][j] {
		return false
	}
	if instack[i][j] {
		instack[i][j] = false
		return true
	}
	instack[i][j] = true
	d := dir[i][j]
	for ni, nj := i+dx[d], j+dy[d]; ; ni, nj = ni+dx[d], nj+dy[d] {
		if ni < 0 || ni >= n || nj < 0 || nj >= n {
			break
		}
		if dfs(ni, nj) {
			dir[ni][nj] = d
			if instack[i][j] {
				instack[i][j] = false
				return true
			} else {
				return dfs(i, j)
			}
		}
	}
	used[i][j] = true
	instack[i][j] = false
	tmp := i
	if dir[i][j]&1 != 0 {
		tmp = j
	}
	fmt.Print(string("LURD"[dir[i][j]]), tmp+1, "\n")
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &u[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i])
	}
	match()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j)
		}
	}
}
