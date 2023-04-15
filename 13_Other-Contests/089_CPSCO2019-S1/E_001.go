package main

import (
	"bufio"
	"fmt"
	"os"
)

const K = 21
const KM = 1 << (K - 1)

var Z, ind, tmp [900000]int
var dat [1 << K]int
var laz [1 << K]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &Z[i])
	}
	var X [200000]int
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &Z[i+N], &Z[i+N+Q], &X[i])
		Z[i+N+Q+Q] = X[i]
	}
	for i := 0; i < N+Q*3; i++ {
		ind[i] = i
	}

	pairsort(N + Q*3)
	j := 0
	mae := Z[0]
	const b = (1 << 30)
	for i := 0; i < N+Q*3; i++ {
		if mae != Z[i] {
			mae = Z[i]
			j++
		}
		k := ind[i]
		if k < N {
			dat[j+KM] ^= Z[i] ^ b
		} else {
			tmp[k] = j
		}
	}
	for i := KM - 1; i > 0; i-- {
		dat[i] = dat[i<<1] ^ dat[i<<1|1]
	}
	L := tmp[N:]
	R := tmp[N+Q:]
	x := tmp[N+Q+Q:]
	const m = (1 << 30) - 1
	for i := 0; i < Q; i++ {
		p := check(L[i], R[i]+1, 1, 0, KM)
		fmt.Fprintln(out, p&m)
		if ((p >> 30) & 1) != 0 {
			a := x[i] + KM
			c := X[i] ^ b
			for j := 20; j >= 0; j-- {
				a2 := (a >> j)
				if laz[a2] {
					if a2 < KM {
						laz[a2*2] = true
						laz[a2*2+1] = true
					}
					dat[a2] = c
					laz[a2] = false
				} else {
					dat[a2] ^= c
				}
			}
		}
	}
}

func pairsort(N int) {
	const b = 8
	for k := 0; k < 4; k++ {
		var kazu [1 << b]int
		var kazu2 [1 << b]int
		for i := 0; i < N; i++ {
			kazu[(Z[i]>>(k*b))&((1<<b)-1)]++
		}
		for i := 0; i < (1<<b)-1; i++ {
			kazu[i+1] += kazu[i]
		}
		for i := N - 1; i >= 0; i-- {
			kazu[(Z[i]>>(k*b))&((1<<b)-1)]--
			dat[kazu[(Z[i]>>(k*b))&((1<<b)-1)]] = Z[i]
			tmp[kazu[(Z[i]>>(k*b))&((1<<b)-1)]] = ind[i]
		}
		k++
		for i := 0; i < N; i++ {
			kazu2[(dat[i]>>(k*b))&((1<<b)-1)]++
		}
		for i := 0; i < (1<<b)-1; i++ {
			kazu2[i+1] += kazu2[i]
		}
		for i := N - 1; i >= 0; i-- {
			kazu2[(dat[i]>>(k*b))&((1<<b)-1)]--
			Z[kazu2[(dat[i]>>(k*b))&((1<<b)-1)]] = dat[i]
			ind[kazu2[(dat[i]>>(k*b))&((1<<b)-1)]] = tmp[i]
		}
	}
}

func check(L, R, A, l, r int) int {
	if laz[A] {
		return 0
	}
	if r <= L || R <= l {
		return 0
	}
	if L <= l && r <= R {
		laz[A] = true
		return dat[A]
	}
	p := check(L, R, A*2, l, (l+r)/2) ^ check(L, R, A*2+1, (l+r)/2, r)
	dat[A] = dat[A] ^ p
	return p
}
