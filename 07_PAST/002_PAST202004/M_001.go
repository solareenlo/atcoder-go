package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var D, L, N int
	fmt.Fscan(in, &D, &L, &N)
	C := make([]int, D)
	for i := range C {
		fmt.Fscan(in, &C[i])
	}

	kosuu := [100001]int{}
	ind := make([][]int, 100001)
	shokuji := make([][]int, 100001)

	for i := 0; i < D; i++ {
		kosuu[C[i]]++
		if len(ind[C[i]]) > 0 {
			in := ind[C[i]][len(ind[C[i]])-1]
			s := shokuji[C[i]][len(shokuji[C[i]])-1]
			ind[C[i]] = append(ind[C[i]], i)
			shokuji[C[i]] = append(shokuji[C[i]], s+(i-in-1)/L+1)
		} else {
			ind[C[i]] = append(ind[C[i]], i)
			shokuji[C[i]] = append(shokuji[C[i]], 0)
		}
	}

	isshuu := [100001]int{}
	for i := 0; i < D; i++ {
		in := ind[C[i]][len(ind[C[i]])-1]
		s := shokuji[C[i]][len(shokuji[C[i]])-1]
		ind[C[i]] = append(ind[C[i]], i+D)
		shokuji[C[i]] = append(shokuji[C[i]], s+(D+i-in-1)/L+1)
		if isshuu[C[i]] == 0 {
			isshuu[C[i]] = shokuji[C[i]][len(shokuji[C[i]])-1] - shokuji[C[i]][0]
		}
	}

	for i := 0; i < N; i++ {
		var K, F, T int
		fmt.Fscan(in, &K, &F, &T)
		if kosuu[K] == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		F--
		idx := lowerBound(ind[K], F)
		ans := 0
		t := (ind[K][idx] - F + L - 1) / L
		T -= t
		if T > 0 {
			tmp := T / isshuu[K]
			T -= tmp * isshuu[K]
			ans += tmp * kosuu[K]
			idx2 := upperBound(shokuji[K][idx:], shokuji[K][idx]+T-1)
			ans += idx2
		}
		fmt.Fprintln(out, ans)
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
