package main

import (
	"bufio"
	"fmt"
	"os"
)

func sort(N int, A []int) {
	const b = 8
	var tmp [1000010]int
	for k := 0; k < 8; k++ {
		var num, num2 [1 << b]int
		for i := 0; i < N; i++ {
			num[(A[i]>>(k*b))&((1<<b)-1)]++
		}
		for i := 0; i < (1<<b)-1; i++ {
			num[i+1] += num[i]
		}
		for i := N - 1; i >= 0; i-- {
			num[(A[i]>>(k*b))&((1<<b)-1)]--
			tmp[num[(A[i]>>(k*b))&((1<<b)-1)]] = A[i]
		}
		k++
		for i := 0; i < N; i++ {
			num2[(tmp[i]>>(k*b))&((1<<b)-1)]++
		}
		for i := 0; i < (1<<b)-1; i++ {
			num2[i+1] += num2[i]
		}
		for i := N - 1; i >= 0; i-- {
			num2[(tmp[i]>>(k*b))&((1<<b)-1)]--
			A[num2[(tmp[i]>>(k*b))&((1<<b)-1)]] = tmp[i]
		}
	}
}

const bm1 = 200010

var BIT1, BIT2 [bm1]int

func add1(A int) {
	B := A
	for A < bm1 {
		BIT1[A] = max(BIT1[A], B)
		A += A & -A
	}
	A = B
	for A > 0 {
		BIT2[A] = min(BIT2[A], B)
		A -= A & -A
	}
}

func query11(A int) int {
	res := 0
	for A > 0 {
		res = max(BIT1[A], res)
		A -= A & -A
	}
	return res
}

func query12(A int) int {
	res := bm1
	for A < bm1 {
		res = min(BIT2[A], res)
		A += A & -A
	}
	return res
}

func del1(A int) {
	pre := query11(A - 1)
	nxt := query12(A + 1)
	B := A
	for A < bm1 && BIT1[A] == B {
		BIT1[A] = 0
		A += A & -A
	}
	if pre != 0 {
		A = pre
		for A < bm1 {
			BIT1[A] = max(BIT1[A], pre)
			A += A & -A
		}
	}
	A = B
	for A > 0 && BIT2[A] == B {
		BIT2[A] = bm1
		A -= A & -A
	}
	if nxt != bm1 {
		A = nxt
		for A > 0 {
			BIT2[A] = min(BIT2[A], nxt)
			A -= A & -A
		}
	}
}

var QD []int
var fu [1000000]bool
var q int

func makeQ(A int) {
	QD[q] = (abs(A) << 20) + q
	if A < 0 {
		fu[q] = true
	}
	q++
}

const bm2 = 1000010

var BIT3 [bm2]int

func add2(A int) {
	B := A
	for A > 0 {
		BIT3[A] = min(BIT3[A], B)
		A -= A & -A
	}
}

func query2(A int) int {
	res := bm2
	for A < bm2 {
		res = min(BIT3[A], res)
		A += A & -A
	}
	return res
}

func del2(A int) {
	nxt := query2(A + 1)
	B := A
	for A > 0 && BIT3[A] == B {
		BIT3[A] = bm2
		A -= A & -A
	}
	if nxt != bm2 {
		A = nxt
		for A > 0 {
			BIT3[A] = min(BIT3[A], nxt)
			A -= A & -A
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, L, X1, X2 int
	fmt.Fscan(in, &N, &L, &X1, &X2)

	var H [200001]int
	tmp2 := make([]int, 1000010)
	H[0] = min(X1, L-X1)
	H[1] = min(X2, L-X2)
	tmp2[0] = X1 << 20
	tmp2[1] = (X2 << 20) + 1
	for i := 1; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		H[i+1] = min(a, L-a)
		tmp2[i+1] = (a << 20) + i + 1
	}
	sort(N+1, tmp2)

	var I [200001]int
	k := 0
	nxt := -1
	for i := 0; i < N+1; i++ {
		if nxt != tmp2[i]>>20 {
			k++
			nxt = tmp2[i] >> 20
		}
		I[tmp2[i]&((1<<20)-1)] = k
	}

	var H2 [200001]bool
	var tmp [200010]int
	H2[0] = true
	H2[1] = true
	tmp[I[0]] = 1
	tmp[I[1]] = 1
	for i := 2; i <= N; i++ {
		if tmp[I[i]] != 0 {
			tmp[I[i]] = 0
			H2[i] = false
		} else {
			tmp[I[i]] = 1
			H2[i] = true
		}
	}

	tmp2[0] = min(X1, L-X1) << 20
	tmp2[1] = (min(X2, L-X2) << 20) + 1
	for i := 1; i < N; i++ {
		tmp2[i+1] = (H[i+1] << 20) + i + 1
	}

	sort(N+1, tmp2)
	k = 1
	var HI [200010]int
	nxt = -1
	for i := 0; i < N+1; i++ {
		if nxt != tmp2[i]>>20 {
			k++
			nxt = tmp2[i] >> 20
			HI[k] = nxt
		}
		I[tmp2[i]&((1<<20)-1)] = k
	}

	for i := 0; i < bm1; i++ {
		BIT2[i] = bm1
	}
	add1(I[0])
	add1(I[1])

	for i := 0; i < N+10; i++ {
		tmp[i] = 0
	}
	tmp[I[0]]++
	tmp[I[1]]++

	QD = make([]int, 1000000)
	makeQ(L - H[0] - H[1])
	fmt.Println(L - H[0] - H[1])

	Q := make([]int, 0)

	for i := 2; i <= N; i++ {
		if H2[i] {
			a := I[i]
			var k1, k2, k3, k4 int
			if tmp[a] != 0 {
				k3 = a
				k4 = query12(k3 + 1)
				k2 = query11(a - 1)
				if k2 == 0 {
					k1 = 0
				} else if tmp[k2] == 2 {
					k1 = k2
				} else {
					k1 = query11(k2 - 1)
				}
				tmp[a] = 2
			} else {
				k3 = query12(a)
				if k3 == bm1 {
					k4 = bm1
				} else if tmp[k3] == 2 {
					k4 = k3
				} else {
					k4 = query12(k3 + 1)
				}
				k2 = query11(a)
				if k2 == 0 {
					k1 = 0
				} else if tmp[k2] == 2 {
					k1 = k2
				} else {
					k1 = query11(k2 - 1)
				}
				add1(a)
				tmp[a] = 1
			}
			a = HI[a]
			if k1 == 0 {
				k1 = -1
			} else {
				k1 = HI[k1]
			}
			if k2 == 0 {
				k2 = -1
			} else {
				k2 = HI[k2]
			}
			if k3 == bm1 {
				k3 = -1
			} else {
				k3 = HI[k3]
			}
			if k4 == bm1 {
				k4 = -1
			} else {
				k4 = HI[k4]
			}
			if k2 == -1 {
				makeQ(k4 - a)
			} else if k3 == -1 {
				makeQ(-(L - k1 - k2))
				makeQ(a - k1)
				makeQ(L - a - k2)
			} else if k1 == -1 && k4 == -1 {
				makeQ(-(L - k2 - k3))
				makeQ(k3 - k2)
				makeQ(L - a - k3)
			} else if k1 == -1 {
				makeQ(-(k4 - k2))
				makeQ(k3 - k2)
				makeQ(k4 - a)
			} else if k4 == -1 {
				makeQ(-(L - k3 - k2))
				makeQ(-(k3 - k1))
				makeQ(a - k1)
				makeQ(k3 - k2)
				makeQ(L - a - k3)
			} else {
				makeQ(-(k4 - k2))
				makeQ(-(k3 - k1))
				makeQ(a - k1)
				makeQ(k3 - k2)
				makeQ(k4 - a)
			}
		} else {
			a := I[i]
			var k1, k2, k3, k4 int
			if tmp[a] == 2 {
				k3 = a
				k4 = query12(k3 + 1)
				k2 = query11(a - 1)
				if k2 == 0 {
					k1 = 0
				} else if tmp[k2] == 2 {
					k1 = k2
				} else {
					k1 = query11(k2 - 1)
				}
				tmp[a] = 1
			} else {
				k3 = query12(a + 1)
				if k3 == bm1 {
					k4 = bm1
				} else if tmp[k3] == 2 {
					k4 = k3
				} else {
					k4 = query12(k3 + 1)
				}
				k2 = query11(a - 1)
				if k2 == 0 {
					k1 = 0
				} else if tmp[k2] == 2 {
					k1 = k2
				} else {
					k1 = query11(k2 - 1)
				}
				del1(a)
				tmp[a] = 0
			}
			a = HI[a]
			if k1 == 0 {
				k1 = -1
			} else {
				k1 = HI[k1]
			}
			if k2 == 0 {
				k2 = -1
			} else {
				k2 = HI[k2]
			}
			if k3 == bm1 {
				k3 = -1
			} else {
				k3 = HI[k3]
			}
			if k4 == bm1 {
				k4 = -1
			} else {
				k4 = HI[k4]
			}
			if k2 == -1 {
				makeQ(-(k4 - a))
			} else if k3 == -1 {
				makeQ(L - k1 - k2)
				makeQ(-(a - k1))
				makeQ(-(L - a - k2))
			} else if k1 == -1 && k4 == -1 {
				makeQ(L - k2 - k3)
				makeQ(-(k3 - k2))
				makeQ(-(L - a - k3))
			} else if k1 == -1 {
				makeQ(k4 - k2)
				makeQ(-(k3 - k2))
				makeQ(-(k4 - a))
			} else if k4 == -1 {
				makeQ(L - k3 - k2)
				makeQ(k3 - k1)
				makeQ(-(a - k1))
				makeQ(-(k3 - k2))
				makeQ(-(L - a - k3))
			} else {
				makeQ(k4 - k2)
				makeQ(k3 - k1)
				makeQ(-(a - k1))
				makeQ(-(k3 - k2))
				makeQ(-(k4 - a))
			}
		}
		Q = append(Q, q-1)
	}

	sort(q, QD)
	kari := 0
	nxt = -1
	var qi, qd [1000010]int
	for i := 0; i < q; i++ {
		if nxt != QD[i]>>20 {
			nxt = QD[i] >> 20
			kari++
			qd[kari] = nxt
		}
		qi[QD[i]&((1<<20)-1)] = kari
	}

	for i := 0; i < bm2; i++ {
		BIT3[i] = bm2
	}
	var kazu [1000000]int
	kari = 0
	for i := 0; i < q; i++ {
		if fu[i] {
			kazu[qi[i]]--
			if kazu[qi[i]] == 0 {
				del2(qi[i])
			}
		} else {
			if kazu[qi[i]] == 0 {
				add2(qi[i])
			}
			kazu[qi[i]]++
		}
		if kari < len(Q) && Q[kari] == i {
			fmt.Println(qd[query2(1)])
			kari++
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
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
