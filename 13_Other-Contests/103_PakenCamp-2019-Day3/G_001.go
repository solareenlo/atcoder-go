package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const big = int(2.19e15) + 1

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const siz = 524288

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)

	atu := make([]int, 0)
	atu = append(atu, -big)
	atu = append(atu, big)
	house := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &house[i])
		atu = append(atu, house[i])
	}

	var m int
	fmt.Fscan(in, &m)

	santa := make([]int, m+2)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &santa[i])
		atu = append(atu, santa[i])
	}
	santa[m] = -big
	santa[m+1] = big
	var Q int
	fmt.Fscan(in, &Q)
	query := make([]pair, 0)
	mH := make([]int, n)
	copy(mH, house)
	mS := make([]int, m+2)
	copy(mS, santa)
	for i := 0; i < Q; i++ {
		var T, C, D int
		fmt.Fscan(in, &T, &C, &D)
		atu = append(atu, D)
		C--
		if T == 1 {
			query = append(query, pair{0, house[C]})
			house[C] = D
			query = append(query, pair{1, house[C]})
		}
		if T == 2 {
			query = append(query, pair{0, santa[C]})
			santa[C] = D
			query = append(query, pair{2, santa[C]})
		}
	}
	copy(house, mH)
	copy(santa, mS)
	sort.Ints(atu)
	atu = unique(atu)
	var iee, san mno
	var seg [siz + siz]mno
	for i := siz; i < siz+siz; i++ {
		seg[i].zyo = 0
	}
	iee.zyo = 1
	iee.L_Dh = 0
	iee.L_Mh = 0
	iee.L_Uh = 0
	iee.R_Dh = 0
	iee.R_Mh = 0
	iee.R_Uh = 0

	san.zyo = 2
	san.L_Dh = 0
	san.L_Mh = 0
	san.L_Uh = 0
	san.R_Dh = 0
	san.R_Mh = 0
	san.R_Uh = 0
	san.LL = 0
	san.LR = big
	san.RL = big
	san.RR = 0
	for i := 0; i < n; i++ {
		X := house[i]
		dco := lowerBound(atu, X) + siz
		seg[dco] = iee
		seg[dco].Lx = X
		seg[dco].Rx = X
		seg[dco].Lsa = X
		seg[dco].Rsa = X
	}
	for i := 0; i < m+2; i++ {
		X := santa[i]
		dco := lowerBound(atu, X) + siz
		seg[dco] = san
		seg[dco].Lx = X
		seg[dco].Rx = X
		seg[dco].Lsa = X
		seg[dco].Rsa = X
	}
	for i := siz - 1; i > 0; i-- {
		seg[i] = culc(seg[i+i], seg[i+i+1])
	}
	fmt.Fprintln(out, min(seg[1].LL, seg[1].LR, seg[1].RL, seg[1].RR))
	for i := 0; i < Q+Q; i++ {
		que := query[i].x
		X := query[i].y
		dco := lowerBound(atu, X) + siz
		if que == 0 {
			seg[dco].zyo = 0
		}
		if que == 1 {
			seg[dco] = iee
			seg[dco].Lx = X
			seg[dco].Rx = X
			seg[dco].Lsa = X
			seg[dco].Rsa = X
		}
		if que == 2 {
			seg[dco] = san
			seg[dco].Lx = X
			seg[dco].Rx = X
			seg[dco].Lsa = X
			seg[dco].Rsa = X
		}
		for dco > 1 {
			dco /= 2
			seg[dco] = culc(seg[dco+dco], seg[dco+dco+1])
		}
		if i%2 == 1 {
			fmt.Fprintln(out, min(seg[1].LL, seg[1].LR, seg[1].RL, seg[1].RR))
		}
	}
}

type mno struct {
	zyo              int
	Lx, Rx, Lsa, Rsa int
	L_Dh, L_Mh, L_Uh int
	R_Dh, R_Mh, R_Uh int
	LL, LR, RL, RR   int
}

func culc(Lin, Rin mno) mno {
	if Lin.zyo == 0 {
		return Rin
	}
	if Rin.zyo == 0 {
		return Lin
	}
	var ans mno
	ans.Lx = Lin.Lx
	ans.Rx = Rin.Rx
	ans.Lsa = Lin.Lsa
	ans.L_Dh = Lin.L_Dh
	ans.L_Mh = Lin.L_Mh
	ans.L_Uh = Lin.L_Uh
	ans.Rsa = Rin.Rsa
	ans.R_Dh = Rin.R_Dh
	ans.R_Mh = Rin.R_Mh
	ans.R_Uh = Rin.R_Uh
	if Lin.zyo == 1 {
		ans.Lsa = Rin.Lsa
		ans.L_Dh = max(Lin.L_Dh-(Rin.Lsa-Lin.Rx), (Rin.Lx-Lin.Rx)-(Rin.Lsa-Rin.Lx), Rin.L_Dh)
		ans.L_Mh = max(Lin.L_Mh, (Rin.Lx - Lin.Rx), Rin.L_Mh)
		ans.L_Uh = max(Lin.L_Uh+(Rin.Lsa-Lin.Rx), (Rin.Lx-Lin.Rx)*2+(Rin.Lsa-Rin.Lx), Rin.L_Uh)
	}
	if Rin.zyo == 1 {
		ans.Rsa = Lin.Rsa
		ans.R_Dh = max(Rin.R_Dh-(Rin.Lx-Lin.Rsa), (Rin.Lx-Lin.Rx)-(Lin.Rx-Lin.Rsa), Lin.R_Dh)
		ans.R_Mh = max(Rin.R_Mh, (Rin.Lx - Lin.Rx), Lin.R_Mh)
		ans.R_Uh = max(Rin.R_Uh+(Rin.Lx-Lin.Rsa), (Rin.Lx-Lin.Rx)*2+(Lin.Rx-Lin.Rsa), Lin.R_Uh)
	}
	if Lin.zyo == 1 && Rin.zyo == 1 {
		ans.zyo = 1
		return ans
	}
	ans.zyo = 2
	if Lin.zyo == 2 && Rin.zyo == 2 {
		var LdL, LdR, RdL, RdR int
		LdL = min((Rin.Lsa-Lin.Rsa)*2-Rin.L_Uh, Rin.Lsa-Lin.Rsa*2-Rin.Lx+Lin.Rx*2, (Rin.Lsa-Lin.Rsa)-Lin.R_Dh)

		RdR = min((Rin.Lsa-Lin.Rsa)*2-Lin.R_Uh, Rin.Lsa*2-Lin.Rsa-Rin.Lx*2+Lin.Rx, (Rin.Lsa-Lin.Rsa)-Rin.L_Dh)

		RdL = min((Rin.Lsa-Lin.Rsa)-Lin.R_Mh, (Rin.Lsa-Lin.Rsa)-(Rin.Lx-Lin.Rx), (Rin.Lsa-Lin.Rsa)-Rin.L_Mh)
		LdR = RdL * 2

		ans.LL = min(Lin.LL+LdL+Rin.LL, Lin.LR+RdL+Rin.LL, Lin.LL+LdR+Rin.RL, Lin.LR+RdR+Rin.RL)
		ans.RL = min(Lin.RL+LdL+Rin.LL, Lin.RR+RdL+Rin.LL, Lin.RL+LdR+Rin.RL, Lin.RR+RdR+Rin.RL)
		ans.LR = min(Lin.LL+LdL+Rin.LR, Lin.LR+RdL+Rin.LR, Lin.LL+LdR+Rin.RR, Lin.LR+RdR+Rin.RR)
		ans.RR = min(Lin.RL+LdL+Rin.LR, Lin.RR+RdL+Rin.LR, Lin.RL+LdR+Rin.RR, Lin.RR+RdR+Rin.RR)
	}
	if Lin.zyo == 1 && Rin.zyo == 2 {
		ans.LL = Rin.LL
		ans.LR = Rin.LR
		ans.RL = Rin.RL
		ans.RR = Rin.RR
	}
	if Lin.zyo == 2 && Rin.zyo == 1 {
		ans.LL = Lin.LL
		ans.LR = Lin.LR
		ans.RL = Lin.RL
		ans.RR = Lin.RR
	}
	return ans
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
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
