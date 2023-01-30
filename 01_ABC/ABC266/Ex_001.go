package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 1e5 + 10

type snuke struct {
	a, b, y, siz, t int
	dp              int
}

var sn [N]snuke
var tot int
var lis [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	num := 0
	for i := 1; i <= n; i++ {
		var t, x, y, a int
		fmt.Fscan(in, &t, &x, &y, &a)
		if x+y-t <= 0 {
			num++
			sn[num].a = x + y - t
			sn[num].b = x - y + t
			sn[num].y = y
			sn[num].siz = a
			sn[num].dp = a
			sn[num].t = t
			tot++
			lis[tot] = sn[num].b
		}
	}
	if num == 0 {
		fmt.Println(0)
		return
	}
	tmp := lis[1 : tot+1]
	sort.Ints(tmp)
	tmp1 := lis[1 : tot+1]
	tot = len(unique(tmp1))
	for i := 1; i <= num; i++ {
		tmp := lis[1 : tot+1]
		sn[i].b = lowerBound(tmp, sn[i].b) + 1
	}
	CDQ(1, num)
	ans := 0
	for i := 1; i <= num; i++ {
		ans = Max(ans, sn[i].dp)
	}
	fmt.Println(ans)
}

var bit bitTree

type bitTree struct {
	t [N]int
}

func (b *bitTree) clear(i int) {
	for i <= tot {
		b.t[i] = 0
		i += i & -i
	}
}

func (b *bitTree) update(i, d int) {
	for i <= tot {
		b.t[i] = Max(b.t[i], d)
		i += i & -i
	}
}

func (b *bitTree) query(i int) int {
	mx := 0
	for i > 0 {
		mx = Max(b.t[i], mx)
		i -= i & -i
	}
	return mx
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CDQ(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	tmp := sn[l : r+1]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].y == tmp[j].y {
			if tmp[i].t == tmp[j].t {
				return tmp[i].b < tmp[j].b
			}
			return tmp[i].t < tmp[j].t
		}
		return tmp[i].y < tmp[j].y

	})
	CDQ(l, mid)
	tmp1 := sn[l : mid+1]
	sort.Slice(tmp1, func(i, j int) bool {
		if tmp1[i].a == tmp1[j].a {
			return tmp1[i].b < tmp1[j].b
		}
		return tmp1[i].a < tmp1[j].a
	})
	tmp2 := sn[mid+1 : r+1]
	sort.Slice(tmp2, func(i, j int) bool {
		if tmp2[i].a == tmp2[j].a {
			return tmp2[i].b < tmp2[j].b
		}
		return tmp2[i].a < tmp2[j].a
	})
	j := mid
	for i := r; i >= mid+1; i-- {
		for j >= l && sn[j].a >= sn[i].a {
			bit.update(sn[j].b, sn[j].dp)
			j--
		}
		sn[i].dp = Max(sn[i].dp, bit.query(sn[i].b)+sn[i].siz)
	}
	for i := j + 1; i <= mid; i++ {
		bit.clear(sn[i].b)
	}
	CDQ(mid+1, r)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
