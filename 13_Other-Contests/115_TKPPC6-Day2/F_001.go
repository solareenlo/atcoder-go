package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100005

	type pair struct {
		x, y int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	if n*(n-1)/2 < m {
		fmt.Println("No")
		return
	}
	var d [N]int
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &d[i])
	}
	mn := 0x3f3f3f3f
	for i := 1; i < n+1; i++ {
		mn = min(mn, d[i])
	}
	if d[1] != d[n] || mn < d[1] {
		fmt.Println("No")
		return
	}
	cnt := 0
	c1 := 0
	for i := 1; i < n+1; i++ {
		if d[i] != mn && d[i]%2 == d[1]%2 {
			c1++
		}
		if d[i] == mn {
			cnt++
		}
	}
	ans := cnt - 1 + c1 + n - cnt - c1
	if n-cnt-c1 != 0 {
		ans++
	}
	if ans > m {
		fmt.Println("No")
		return
	}
	for i := 1; i < n+1; i++ {
		if d[i] != mn {
			tmp := 0
			if mn == 0 && (d[i]&1) != 0 {
				tmp = 1
			}
			if d[i] == 1 || tmp != 0 {
				fmt.Println("No")
				return
			}
		}
	}

	fmt.Println("Yes")
	lst := 1
	vis := make(map[pair]bool)
	for i := 2; i < n+1; i++ {
		if d[i] == mn {
			vis[pair{lst, i}] = true
			if i != n {
				fmt.Println(lst, i, 0)
			} else {
				fmt.Println(lst, i, mn)
			}
			lst = i
		}
	}

	pos := 0
	for i := 1; i < n+1; i++ {
		if d[i] != mn {
			if d[i]%2 == d[1]%2 {
				vis[pair{1, i}] = true
				fmt.Println(1, i, (d[i]-d[1])/2)
			} else {
				if pos == 0 {
					pos = i
				} else if d[pos] > d[i] {
					pos = i
				}
			}
		}
	}
	if pos != 0 {
		vis[pair{1, pos}] = true
		vis[pair{pos, n}] = true
		fmt.Println(1, pos, d[pos]/2)
		fmt.Println(pos, n, (d[pos]+1)/2)
	}
	for i := 1; i < n+1; i++ {
		if d[i] != mn {
			if d[i]%2 != d[1]%2 && i != pos {
				if i > pos {
					vis[pair{pos, i}] = true
				} else {
					vis[pair{i, pos}] = true
				}
				fmt.Println(min(i, pos), max(i, pos), (d[i]-d[pos])/2)
			}
		}
	}

	x, y := 1, 2
	for i := 1; i < (m-ans)+1; i++ {
		if (vis[pair{x, y}]) {
			if y == n {
				x++
				y = x + 1
			} else {
				y++
			}
			i--
		} else {
			fmt.Println(x, y, 1000000000)
			if y == n {
				x++
				y = x + 1
			} else {
				y++
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
