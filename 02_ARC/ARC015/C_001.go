package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	dic := map[string]int{}
	var f, t string
	dp := [201][201]float64{}
	for i := 0; i < n; i++ {
		var cost float64
		fmt.Scan(&f, &cost, &t)
		if _, ok := dic[f]; !ok {
			dic[f] = len(dic)
		}
		if _, ok := dic[t]; !ok {
			dic[t] = len(dic)
		}
		dp[dic[f]][dic[t]] = cost
		dp[dic[t]][dic[f]] = 1.0 / cost
	}

	m := len(dic)
	for i := 0; i < m; i++ {
		dp[i][i] = 1
	}
	for k := 0; k < m; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if dp[i][j] == 0 {
					dp[i][j] = dp[i][k] * dp[k][j]
				}
			}
		}
	}

	mx := -1.0
	for aKey, aVal := range dic {
		for bKey, bVal := range dic {
			if mx < dp[aVal][bVal] {
				mx = dp[aVal][bVal]
				f = aKey
				t = bKey
			}
		}
	}
	fmt.Print("1", f, "=", int(mx+0.5), t)
	fmt.Println()
}
