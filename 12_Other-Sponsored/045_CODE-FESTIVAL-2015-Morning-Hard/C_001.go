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

	var n, k int
	fmt.Fscan(in, &n, &k)
	v := make([]int, n)
	rv := make([]int, n)
	f := make([]bool, n+1)
	f[0] = true
	f[n] = true
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
		v[i]--
		rv[v[i]] = i
	}
	x := 0
	for x < n {
		if k == 0 {
			break
		}
		i := rv[x]
		if !f[i] {
			f[i] = true
			k--
		}
		if k == 0 {
			break
		}
		r := i
		for r+1 < n && v[r+1] == v[r]+1 {
			r++
		}
		if !f[r+1] {
			if k == 1 {
				srt := make([]int, 0)
				for j := 0; j < n; j++ {
					if f[j] && v[j] >= v[i] {
						srt = append(srt, v[j])
					}
				}
				sort.Ints(srt)
				sort.Sort(sort.Reverse(sort.IntSlice(srt)))
				done := false
				d := make([]bool, n+1)
				for j := v[i]; j < n; j++ {
					if done {
						break
					}
					if !d[rv[j]] && !f[rv[j]] {
						f[rv[j]] = true
						break
					}
					if d[rv[j]] {
						continue
					}
					srt = srt[:len(srt)-1]
					if len(srt) == 0 {
						break
					}
					for k := rv[j] + 1; k < n; k++ {
						if f[k] {
							break
						}
						if v[k] > srt[len(srt)-1] {
							f[k] = true
							done = true
							break
						}
						d[k] = true
					}
				}
				break
			}
			f[r+1] = true
			k--
		}
		x = v[r] + 1
	}
	al := make([][]int, 0)
	for i := 0; i < n; i++ {
		if f[i] {
			tmp := make([]int, 0)
			tmp = append(tmp, v[i])
			for j := i + 1; j < n; j++ {
				if f[j] {
					break
				}
				tmp = append(tmp, v[j])
			}
			al = append(al, tmp)
		}
	}
	srt := make([]int, 0)
	for i := 0; i < len(al); i++ {
		srt = append(srt, i)
	}
	sort.Slice(srt, func(a, b int) bool {
		return al[srt[a]][0] < al[srt[b]][0]
	})
	for _, i := range srt {
		for _, j := range al[i] {
			fmt.Fprintln(out, j+1)
		}
	}
}
