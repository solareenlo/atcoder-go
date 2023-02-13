package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	ans := 0
	for i := n - 1; i > 4 && a[i]+a[i-1]+a[i-2]+a[i-3]+a[i-4]+a[i-5] > ans; i-- {
		for j := i - 1; j > 0 && a[i]+a[j]+a[j-1]+a[i-1]+a[i-2]+a[i-3] > ans; j-- {
			for k := j - 1; k >= 0 && a[j]+a[k] > a[i] && (a[i]+a[j]+a[k])+a[i-1]+a[i-2]+a[i-3] > ans; k-- {
				l := a[i] + a[j] + a[k]
				for x := i - 1; x > 1 && l+a[x]+a[x-1]+a[x-2] > ans; x-- {
					if x != j && x != k {
						for y := x - 1; y > 0 && l+a[x]+a[y]+a[y-1] > ans; y-- {
							if y != j && y != k {
								for z := y - 1; z >= 0 && a[y]+a[z] > a[x] && l+a[x]+a[y]+a[z] > ans; z-- {
									if z != j && z != k {
										ans = l + a[x] + a[y] + a[z]
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
