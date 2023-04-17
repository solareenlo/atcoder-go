package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var N, M, SA, SB int
var A, B [20]int
var VA, VB []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &M)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		SA += A[i]
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &B[i])
		SB += B[i]
	}

	sagasuA(0, 0)
	sagasuB(0, 0)

	sort.Ints(VA)
	sort.Ints(VB)

	maxi := 0.0
	k, k2 := 0, 0
	for i := 0; i < len(VA); i++ {
		for k < len(VB) && VA[i] >= VB[k] {
			k++
		}
		for k2 < len(VB) && SB-SA+VA[i] >= VB[k2] {
			k2++
		}
		kari := float64(k - k2)
		maxi = math.Max(maxi, float64(kari)/float64(len(VB)))
	}
	fmt.Println(maxi)
}

func sagasuA(p, kei int) {
	if kei+kei >= SA {
		return
	}
	if p < N {
		sagasuA(p+1, kei)
		sagasuA(p+1, kei+A[p])
	} else if kei != 0 {
		VA = append(VA, SA-kei)
	}
}

func sagasuB(p, kei int) {
	if kei+kei >= SB {
		return
	}
	if p < M {
		sagasuB(p+1, kei)
		sagasuB(p+1, kei+B[p])
	} else if kei != 0 {
		VB = append(VB, SB-kei)
	}
}
