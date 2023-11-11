package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	x := make([]int, n+1)
	y := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &y[i])
	}
	var q *list.List
	q = list.New()
	dif, ans, sum := 0, 0, 0
	for i := 1; i <= n; i++ {
		fr, ba := 0, 0
		for q.Len() > 0 {
			cur := q.Front().Value.(pair).x
			cnt := q.Front().Value.(pair).y
			cur += dif
			if cur >= -x[i] {
				break
			}
			q.Remove(q.Front())
			fr += cnt
			sum = (sum + (cur+x[i])%mod*cnt%mod) % mod
		}
		for q.Len() > 0 {
			cur := q.Back().Value.(pair).x
			cnt := q.Back().Value.(pair).y
			cur += dif
			if cur <= y[i] {
				break
			}
			q.Remove(q.Back())
			ba += cnt
			sum = (sum + ((-cur+y[i])%mod+mod)%mod*cnt%mod) % mod
		}
		sum += (x[i] + y[i]) * (i - 1) * 2 % mod
		sum %= mod
		sum += (x[i] + y[i]) * 3 % mod
		sum %= mod
		dif += x[i] - y[i]
		q.PushFront(pair{-y[i] - dif, fr + 1})
		q.PushBack(pair{x[i] - dif, ba + 1})
		ans = (ans + sum) % mod
	}
	ans = (ans * 2) % mod
	for i := 1; i <= n; i++ {
		ans = ((ans-x[i]*3-y[i]*3)%mod + mod) % mod
	}
	fmt.Println(ans)
}
