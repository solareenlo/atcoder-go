package main

import (
	"fmt"
	"sort"
)

const INF = int(5e8)

var g [6000][]int
var ar [100005]int
var perm [][]int
var dp [25][6000]int
var Type [100005]int
var swapPos [10][10][]int

func main() {
	var w, h int
	fmt.Scan(&w, &h)

	K := 1
	for i := 0; i < w; i++ {
		K *= i + 1
	}

	L := w*(w-1)/2 + 1

	trg := make([]int, w)
	for i := 0; i < h; i++ {
		fmt.Scan(&ar[i])
	}
	for i := 0; i < w; i++ {
		fmt.Scan(&trg[i])
	}

	tmp := make([]int, w)
	for i := 0; i < w; i++ {
		tmp[i] = i
	}
	for {
		perm = append(perm, make([]int, len(tmp)))
		copy(perm[len(perm)-1], tmp)
		if !nextPermutation(tmp) {
			break
		}
	}

	cur := make([]int, w)
	for i := 0; i < w; i++ {
		cur[i] = i
	}

	Type[h] = 0
	for i := h - 1; i >= 0; i-- {
		a := ar[i]
		cur[a], cur[a+1] = cur[a+1], cur[a]
		Type[i] = lowerBoundSlice(perm, cur)
	}

	for i := 0; i < h+1; i++ {
		t := Type[i]
		p := perm[t]

		for j := 0; j < w-1; j++ {
			a, b := p[j], p[j+1]
			if a > b {
				a, b = b, a
			}
			swapPos[a][b] = append(swapPos[a][b], i)
		}
	}
	for j := 0; j < w; j++ {
		for j2 := j + 1; j2 < w; j2++ {
			swapPos[j][j2] = append(swapPos[j][j2], INF)
		}
	}

	rev := make([]int, w)
	for i := 0; i < w; i++ {
		rev[cur[i]] = i
	}

	initState := lowerBoundSlice(perm, rev)

	for i := 0; i < L; i++ {
		for j := 0; j < K; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][initState] = 0

	for i := 0; i < L; i++ {
		for j := 0; j < K; j++ {
			if dp[i][j] < INF {
				if i+1 < L {
					for s := 0; s < w; s++ {
						for t := s + 1; t < w; t++ {
							lb := sort.Search(len(swapPos[s][t]), func(k int) bool {
								return swapPos[s][t][k] >= dp[i][j]
							})
							if swapPos[s][t][lb] == INF {
								continue
							}
							p := make([]int, len(perm[j]))
							copy(p, perm[j])
							p[s], p[t] = p[t], p[s]
							id := lowerBoundSlice(perm, p)
							dp[i+1][id] = min(dp[i+1][id], swapPos[s][t][lb])
						}
					}
				}
			}
		}
	}

	goal := lowerBoundSlice(perm, trg)
	res := INF
	for i := 0; i < L; i++ {
		if dp[i][goal] < INF {
			res = i
			break
		}
	}
	fmt.Println(res)
}

func nextPermutation(a []int) bool {
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := len(a) - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	for l, r := i+1, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return true
}

func lowerBoundSlice(a [][]int, x []int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return !lessThanSlices(a[i], x)
	})
	return idx
}

func lessThanSlices(s1, s2 []int) bool {
	minLength := len(s1)
	if len(s2) < minLength {
		minLength = len(s2)
	}

	for i := 0; i < minLength; i++ {
		if s1[i] < s2[i] {
			return true // s1がs2よりも小さい
		} else if s1[i] > s2[i] {
			return false // s1がs2よりも大きい
		}
	}

	if len(s1) == len(s2) {
		return false // スライスは等しい
	} else if len(s1) < len(s2) {
		return true // s1がs2よりも小さい
	}

	return false // s1がs2よりも大きい
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
