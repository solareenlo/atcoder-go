package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)

var (
	dp        = [20][100][12][12]int{}
	tway      = [12][20]int{}
	way       = [12][20]int{}
	num       = [20]int{}
	n, length int
	ans       int = MaxInt
	vis           = [10]bool{}
)

func dfs(dep, sum, free, u, res int) {
	if dep == n+1 {
		if sum != 0 {
			return
		}
		tmp := 0
		if free == 1 {
			tmp = 1
		}
		if ans > res-tmp {
			length = free
			ans = res - tmp
			for i := 1; i <= length; i++ {
				for j := 1; j <= n; j++ {
					way[i][j] = tway[i][j]
				}
			}
		}
		return
	}
	if free >= 10 {
		return
	}
	if dp[dep][sum][free][u] <= res {
		return
	}
	dp[dep][sum][free][u] = res
	for i := 0; i <= 9; i++ {
		if i > sum {
			break
		}
		if vis[i] {
			las := sum - i
			if las < 10 {
				if u == 0 {
					tway[free+1][dep] = i
					dfs(dep+1, las*10+num[dep+1], free+1, free+1, res+n-dep+2)
					tway[free+1][dep] = 0
				} else {
					tway[u][dep] = i
					dfs(dep+1, las*10+num[dep+1], free, free, res)
					tway[u][dep] = 0
				}
			}
			if u == 0 {
				tway[free+1][dep] = i
				dfs(dep, las, free+1, 0, res+n-dep+2)
				tway[free+1][dep] = 0
			} else {
				tway[u][dep] = i
				dfs(dep, las, free, u-1, res)
				tway[u][dep] = 0
			}
		}
	}
}

func main() {
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = 2139062143
				}
			}
		}
	}
	var str string
	fmt.Scan(&str)
	str = " " + str
	for i := len(str) - 1; i >= 1; i-- {
		vis[str[i]-'0'] = true
	}

	var tar int
	fmt.Scan(&tar)
	for tar != 0 {
		n++
		num[n] = tar % 10
		tar /= 10
	}
	for i := 1; i <= n/2; i++ {
		num[i], num[n-i+1] = num[n-i+1], num[i]
	}
	dfs(1, num[1], 0, 0, 0)
	for i := 1; i < length; i++ {
		fl := false
		for j := 1; j <= n; j++ {
			if !fl && way[i][j] == 0 {
				continue
			}
			fl = true
			fmt.Print(way[i][j])
		}
		fmt.Print("+")
	}
	fl := false
	for j := 1; j <= n; j++ {
		if !fl && way[length][j] == 0 {
			continue
		}
		fl = true
		fmt.Print(way[length][j])
	}
	if length != 1 {
		fmt.Println("=")
	}
}
