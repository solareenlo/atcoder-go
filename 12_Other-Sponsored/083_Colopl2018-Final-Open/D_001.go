package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD = 1000000007

const maxN = 1 << 18

var arr [2 * maxN]int

func inc(a int) {
	for i := a; i <= maxN; i += i & -i {
		arr[i]++
	}
}

func sum(a int) int {
	res := 0
	for i := a; i > 0; i -= i & -i {
		res += arr[i]
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	vp := make([]pair, N)
	for p := range vp {
		fmt.Fscan(in, &vp[p].x, &vp[p].y)
		if vp[p].x > vp[p].y {
			vp[p].x, vp[p].y = vp[p].y, vp[p].x
		}
	}
	sortPair(vp)
	add := 1
	for i := 0; i < N-2; i++ {
		add = (2 * add) % MOD
	}
	res := 0
	for i := len(vp) - 1; i >= 0; i-- {
		res += sum(vp[i].y-1) * add
		res %= MOD
		inc(vp[i].x)
		inc(vp[i].y)
	}
	fmt.Println(res)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
