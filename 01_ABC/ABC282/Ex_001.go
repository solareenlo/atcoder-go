package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 200020

var ls, rs, sumb, a [MAXN]int
var mxs, ans int

func solve(id, L, R int) {
	if id == 0 || L > R {
		return
	}
	solve(ls[id], L, id-1)
	solve(rs[id], id+1, R)
	if id-L < R-id {
		for i := id; i >= L; i-- {
			l := id
			r := R
			res := id - 1
			for l <= r {
				mid := (l + r) / 2
				if sumb[mid]-sumb[i-1] <= mxs-a[id] {
					l = mid + 1
					res = mid
				} else {
					r = mid - 1
				}
			}
			ans += res - id + 1
		}
	} else {
		for i := id; i <= R; i++ {
			l := L
			r := id
			res := id + 1
			for l <= r {
				mid := (l + r) / 2
				if sumb[i]-sumb[mid-1] <= mxs-a[id] {
					r = mid - 1
					res = mid
				} else {
					l = mid + 1
				}
			}
			ans += id - res + 1
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n, &mxs)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		sumb[i] = sumb[i-1] + tmp
	}
	var st [MAXN]int
	stt := 0
	for i := 1; i <= n; i++ {
		for stt != 0 && a[st[stt]] > a[i] {
			stt--
		}
		ls[i] = rs[st[stt]]
		rs[st[stt]] = i
		stt++
		st[stt] = i
	}
	ls[0] = 0
	rs[0] = 0
	solve(st[1], 1, n)
	fmt.Println(ans)
}
