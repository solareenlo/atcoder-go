package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n)

	words := make([]string, n)
	for i := range words {
		fmt.Scan(&words[i])
	}

	fmt.Scan(&m)
	fmt.Println(m)
	grid := make([][]string, m)
	for i := range grid {
		var t string
		fmt.Scan(&t)
		grid[i] = strings.Split(t, "")
	}

	ys := []int{}
	xs := []int{}
	ls := []int{}
	ds := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == "." {
				cur1, cur2 := i, j
				for cur1 < m && grid[cur1][j] != "#" {
					cur1++
				}
				for cur2 < m && grid[i][cur2] != "#" {
					cur2++
				}
				if cur1 != i+1 && (i == 0 || grid[i-1][j] == "#") {
					ys = append(ys, i)
					xs = append(xs, j)
					ls = append(ls, cur1-i)
					ds = append(ds, 0)
				}
				if cur2 != j+1 && (j == 0 || grid[i][j-1] == "#") {
					ys = append(ys, i)
					xs = append(xs, j)
					ls = append(ls, cur2-j)
					ds = append(ds, 1)
				}
			}
		}
	}

	us := make([]int, len(ys))
	var dfs func(int) int
	dfs = func(cur int) int {
		if cur == n {
			return 1
		}
		tmp := make([]byte, 20)
		for i := 0; i < len(ys); i++ {
			if us[i] == 0 && ls[i] == len(words[cur]) {
				if ds[i] == 0 {
					ok := 1
					for k := 0; k < ls[i]; k++ {
						if grid[ys[i]+k][xs[i]] != "." && grid[ys[i]+k][xs[i]][0] != words[cur][k] {
							ok = 0
							break
						}
					}
					if ok == 1 {
						us[i] = 1
						for k := 0; k < ls[i]; k++ {
							tmp[k] = grid[ys[i]+k][xs[i]][0]
							grid[ys[i]+k][xs[i]] = string(words[cur][k])
						}
						if dfs(cur+1) == 1 {
							return 1
						}
						us[i] = 0
						for k := 0; k < ls[i]; k++ {
							grid[ys[i]+k][xs[i]] = string(tmp[k])
						}
					}
				} else {
					ok := 1
					for k := 0; k < ls[i]; k++ {
						if grid[ys[i]][xs[i]+k] != "." && grid[ys[i]][xs[i]+k][0] != words[cur][k] {
							ok = 0
							break
						}
					}
					if ok == 1 {
						us[i] = 1
						for k := 0; k < ls[i]; k++ {
							tmp[k] = grid[ys[i]][xs[i]+k][0]
							grid[ys[i]][xs[i]+k] = string(words[cur][k])
						}
						if dfs(cur+1) == 1 {
							return 1
						}
						us[i] = 0
						for k := 0; k < ls[i]; k++ {
							grid[ys[i]][xs[i]+k] = string(tmp[k])
						}
					}
				}
			}
		}
		return 0
	}
	dfs(0)
	for i := 0; i < m; i++ {
		fmt.Println(strings.Join(grid[i], ""))
	}
}
