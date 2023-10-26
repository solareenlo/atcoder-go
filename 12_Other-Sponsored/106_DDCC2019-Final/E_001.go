package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	type tuple struct {
		x, y, z int
	}

	var d [2222]int
	var num int
	fmt.Fscan(in, &num)
	for i := 0; i < num; i++ {
		var z int
		fmt.Fscan(in, &z)
		d[z]++
	}
	var c [100]int
	for i := 1; i <= 2000; i += 2 {
		if d[i] == 1 && d[i+1] == 1 {
			c[0]++
		} else {
			c[1]++
		}
	}
	s0 := (c[0] + 33) / 34
	s1 := (c[1] + 33) / 34
	nv := 3 + s0 + s1 + 34
	v := make([]tuple, 0)
	v = append(v, tuple{1, 2, 1})
	for i := 0; i < s0; i++ {
		v = append(v, tuple{1, 3 + i, 0})
		v = append(v, tuple{2, 3 + i, 0})
	}
	for i := 0; i < s1; i++ {
		v = append(v, tuple{1, 3 + s0 + i, 0})
	}
	for i := 0; i < 34; i++ {
		v = append(v, tuple{3 + s0 + s1 + i, nv, 0})
	}
	v0 := make([]pair, 0)
	v1 := make([]pair, 0)
	for i := 0; i < s0; i++ {
		for j := 0; j < 34; j++ {
			v0 = append(v0, pair{3 + i, 3 + s0 + s1 + j})
		}
	}
	for i := 0; i < s1; i++ {
		for j := 0; j < 34; j++ {
			v1 = append(v1, pair{3 + s0 + i, 3 + s0 + s1 + j})
		}
	}
	p0, p1 := 0, 0
	for i := 1; i <= 2000; i += 2 {
		if d[i] == 1 && d[i+1] == 1 {
			v = append(v, tuple{v0[p0].x, v0[p0].y, i})
			p0++
		} else if d[i] == 1 {
			v = append(v, tuple{v1[p1].x, v1[p1].y, i})
			p1++
		} else if d[i+1] == 1 {
			v = append(v, tuple{v1[p1].x, v1[p1].y, i + 1})
			p1++
		}
	}
	fmt.Fprintln(out, nv, len(v))
	for i := 0; i < len(v); i++ {
		fmt.Fprintln(out, v[i].x, v[i].y, v[i].z)
	}
}
