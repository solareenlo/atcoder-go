package main

import "fmt"

func main() {
	var vis [1000010]bool
	for i := 2; i <= 1e6; i++ {
		if !vis[i] {
			for j := 2; i*j <= 1e6; j++ {
				vis[i*j] = true
			}
		}
	}
	var n int
	fmt.Scan(&n)
	ans := 0
	for i := 2; i*i <= n; i++ {
		if !vis[i] {
			for j, t := 2, n/i/i; j < i && j*j*j <= t; j++ {
				if !vis[j] {
					for k := j + 1; k < i && j*j*k <= t; k++ {
						if !vis[k] {
							ans++
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
