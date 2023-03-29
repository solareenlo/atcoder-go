package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	first  int
	second int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, i, j int
	var p []pair
	var a, f []int
	var v, w []pair

	fmt.Fscan(in, &n)
	p = make([]pair, n)
	a = make([]int, n)
	f = make([]int, n)

	for i = 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		p[i] = pair{x, i}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].first < p[j].first
	})

	for i = 0; i < n; i++ {
		a[p[i].second] = i
	}

	for i = 0; i < n; i++ {
		x := i
		var z []int
		if f[i] == 1 || a[i] == i {
			continue
		}
		for f[x] == 0 {
			f[x] = 1
			z = append(z, x)
			x = a[x]
		}
		for j = 0; j < len(z)/2; j++ {
			v = append(v, pair{z[j], z[len(z)-j-1]})
		}
		for j = 0; j < (len(z)-1)/2; j++ {
			w = append(w, pair{z[j+1], z[len(z)-j-1]})
		}
	}

	if len(v) == 0 {
		fmt.Fprintln(out, "0")
	} else if len(w) == 0 {
		fmt.Fprintln(out, "1")
		fmt.Fprintf(out, "%d", len(v))
		for i = 0; i < len(v); i++ {
			fmt.Fprintf(out, " %d %d", v[i].first+1, v[i].second+1)
		}
		fmt.Fprintln(out)
	} else {
		fmt.Fprintln(out, "2")
		fmt.Fprintf(out, "%d", len(v))
		for i = 0; i < len(v); i++ {
			fmt.Fprintf(out, " %d %d", v[i].first+1, v[i].second+1)
		}
		fmt.Fprintln(out)
		fmt.Fprintf(out, "%d", len(w))
		for i = 0; i < len(w); i++ {
			fmt.Fprintf(out, " %d %d", w[i].first+1, w[i].second+1)
		}
		fmt.Fprintln(out)
	}
}
