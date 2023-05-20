package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const inf = 1001001001001001001

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = inf
	}

	var sol func(int, int)
	sol = func(l, r int) {
		if r-l == 1 {
			ans[0] = min(ans[0], abs(a[l]-b[l]))
			return
		}
		md := (l + r) / 2
		sol(l, md)
		sol(md, r)
		al := make([]int, 0)
		bl := make([]int, 0)
		ar := make([]int, 0)
		br := make([]int, 0)
		for i := md - 1; i >= l; i-- {
			al = append(al, a[i])
			bl = append(bl, b[i])
		}
		for i := md; i < r; i++ {
			ar = append(ar, a[i])
			br = append(br, b[i])
		}
		for i := 0; i < len(al)-1; i++ {
			al[i+1] = min(al[i+1], al[i])
		}
		for i := 0; i < len(ar)-1; i++ {
			ar[i+1] = min(ar[i+1], ar[i])
		}
		for i := 0; i < len(bl)-1; i++ {
			bl[i+1] = min(bl[i+1], bl[i])
		}
		for i := 0; i < len(br)-1; i++ {
			br[i+1] = min(br[i+1], br[i])
		}
		for tt := 0; tt < 2; tt++ {
			al, ar = ar, al
			bl, br = br, bl
			cur := 0
			tmp := pair{0, -1}
			deq := make([]pair, 0)
			for i := 0; i < len(al)+len(ar)-1; i++ {
				cur = max(cur, i-(len(al))+1)
				for cur < len(ar) && i-cur >= 0 && al[i-cur] <= ar[cur] && bl[i-cur] <= br[cur] {
					cur++
				}
				rng := pair{i - cur + 1, len(al) - 1}
				rng.y = min(rng.y, i)
				if rng.x <= rng.y {
					for j := tmp.y + 1; j < rng.y+1; j++ {
						ad := pair{abs(al[j] - bl[j]), j}
						for len(deq) != 0 && deq[len(deq)-1].x >= ad.x {
							deq = deq[:len(deq)-1]
						}
						deq = append(deq, ad)
					}
					for len(deq) != 0 && deq[0].y < rng.x {
						deq = deq[1:]
					}
					ans[i+1] = min(ans[i+1], deq[0].x)
					tmp = rng
				}
			}
		}
		for tt := 0; tt < 2; tt++ {
			al, ar = ar, al
			bl, br = br, bl
			cur := pair{-1, 0}
			for i := 0; i < len(al)+len(ar)-1; i++ {
				if cur.x == len(al)-1 {
					cur.y++
				} else {
					cur.x++
				}
				for cur.x >= 0 && cur.y < len(br) {
					if bl[cur.x] < br[cur.y] {
						cur.x--
						cur.y++
					} else if cur.y < len(br)-1 && cur.x > 0 && al[cur.x-1] <= ar[cur.y+1] && al[cur.x-1] < br[cur.y+1] {
						cur.x--
						cur.y++
					} else {
						break
					}
				}
				if cur.x >= 0 && cur.y < len(ar) && al[cur.x] <= ar[cur.y] {
					ans[i+1] = min(ans[i+1], abs(al[cur.x]-br[cur.y]))
				}
				if cur.x-1 >= 0 && cur.y+1 < len(ar) && al[cur.x-1] <= ar[cur.y+1] {
					ans[i+1] = min(ans[i+1], abs(al[cur.x-1]-br[cur.y+1]))
				}
			}
		}
	}

	sol(0, n)
	for _, x := range ans {
		fmt.Println(x)
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
