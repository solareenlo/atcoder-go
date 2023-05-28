package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	var SS string
	fmt.Fscan(in, &SS)
	S := make([]int, N)
	for i := 0; i < N; i++ {
		if SS[i] == 'R' {
			S[i] = 0
		} else if SS[i] == 'P' {
			S[i] = 1
		} else {
			S[i] = 2
		}
	}

	ans := 0
	for t := 0; t < 3; t++ {
		a := t
		b := (t + 1) % 3
		c := (t + 2) % 3
		check := make([]int, N)
		for q := 0; q < 2; q++ {
			pos := make([]int, 0)
			posz := make([]int, 0)
			for i := 0; i < N; i++ {
				if S[i] == c {
					pos = append(pos, i)
				}
				if S[i] == a {
					posz = append(posz, i)
				}
			}
			cn := make([]int, 3)
			imo := make([]int, N+3)
			for i := N - 1; i >= 0; i-- {
				if S[i] == b {
					if (cn[a] != 0 && cn[c] != 0) || cn[c] == 0 {
						it := lowerBound(pos, i)
						if it != 0 {
							it--
							imo[pos[it]]++
							imo[pos[it]+1]--
							it++
						}
						it2 := lowerBound(posz, i)
						if it != 0 && it2 != 0 {
							it--
							it2--
							if pos[it] < posz[it2] {
								if it != 0 {
									it--
									imo[0]++
									imo[pos[it]+1]--
								}
							} else {
								imo[0]++
								imo[posz[it2]]--
							}
						}
					}
				}
				cn[S[i]]++
			}
			for i := 1; i < N; i++ {
				imo[i] += imo[i-1]
			}
			imo[N-1]++
			if q == 0 {
				for i := 0; i < N; i++ {
					if S[i] == c && imo[i] != 0 {
						check[i]++
					}
				}
			} else {
				for i := 0; i < N; i++ {
					if S[i] == c && imo[i] != 0 {
						check[N-1-i]++
					}
				}
			}
			S = reverseOrderInt(S)
		}
		for i := 0; i < N; i++ {
			if check[i] == 2 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
