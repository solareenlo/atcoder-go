package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	Len := N - K + 2

	A := make([]int, N)
	s := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		if i > 0 {
			if A[i] != A[i-1] {
				s[i] = s[i-1] + 1
			} else {
				s[i] = s[i-1]
			}
		}
	}

	for l := 0; l+Len <= N; l++ {
		r := l + Len
		if s[r-1]-s[l] < Len-1 || Len%2 == 0 {
			fmt.Println("Yes")
			a := make([]int, 0)
			for i := 0; i < l; i++ {
				a = append(a, 1)
			}
			sl := 0
			sr := accumulate(A[l:r])
			for m := l + 1; m < r; m++ {
				sl += A[m-1]
				sr -= A[m-1]
				if sl%3 != 0 && sr%3 != 0 {
					a = append(a, m-l)
					a = append(a, r-m)
					break
				}
			}
			for i := r; i < N; i++ {
				a = append(a, 1)
			}
			for i := 0; i < K; i++ {
				if i == K-1 {
					fmt.Println(a[i])
				} else {
					fmt.Printf("%d ", a[i])
				}
			}
			return
		}
	}

	fmt.Println("No")
}

func accumulate(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
