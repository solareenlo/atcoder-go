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

	type pair struct {
		x, y int
	}

	var N int
	fmt.Fscan(in, &N)
	In := make([]float64, N)
	for i := range In {
		In[i] = -1
	}
	p := make([]pair, N*2)
	for i := 0; i < N; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		p[i] = pair{s, i}
		p[i+N] = pair{t, i}
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})
	tmp := 1.0
	meat := 0
	ans1 := make([]float64, N)
	ans2 := make([]float64, N)
	for i := 0; i < N*2; i++ {
		ind := p[i].y
		if In[ind] == -1 {
			In[ind] = tmp
			meat++
		} else {
			ans1[ind] = 1 - tmp/In[ind]
			tmp *= float64(meat - 1)
			tmp /= float64(meat)
			ans2[ind] = tmp / In[ind]
			meat--
			if meat == 0 {
				tmp = 1
			}
		}
	}
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, ans1[i], ans2[i])
	}
}
