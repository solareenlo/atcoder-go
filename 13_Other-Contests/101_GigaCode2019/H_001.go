package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type T struct {
	x    string
	y, z int
}

var ans [50000]T
var N int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, K int
	fmt.Fscan(in, &t, &N, &K)

	var S [1024][]string
	if t == 1 {
		for i := 0; i < K; i++ {
			for j := 0; j < N; j++ {
				if (i>>j)%2 == 1 {
					S[i] = append(S[i], "1")
				} else {
					S[i] = append(S[i], "0")
				}
			}
		}
	}
	if t == 2 {
		for i := 0; i < K; i++ {
			var t string
			fmt.Fscan(in, &t)
			S[i] = strings.Split(t, "")
		}
	}

	iter := N
	vec := make([]int, 0)
	for i := 0; i < K; i++ {
		iter = create(S[i], iter)
		vec = append(vec, iter-1)
	}

	if K >= 2 {
		ans[iter] = T{"OR", vec[0], vec[1]}
		iter++
		for i := 2; i < len(vec); i++ {
			ans[iter] = T{"OR", vec[i], iter - 1}
			iter++
		}
	}

	fmt.Println(iter - N)
	for i := N; i < iter; i++ {
		if ans[i].x == "NOT" {
			fmt.Fprintln(out, ans[i].x, ans[i].y+1)
		} else {
			u := (ans[i].y) + 1
			v := (ans[i].z) + 1
			if u > v {
				u, v = v, u
			}
			fmt.Fprintln(out, ans[i].x, u, v)
		}
	}
}

func create(s []string, iter int) int {
	ids := make([]int, 0)
	for i := 0; i < N; i++ {
		if s[i] == "0" {
			ans[iter] = T{"NOT", i, -1}
			ids = append(ids, iter)
			iter++
		} else {
			ids = append(ids, i)
		}
	}
	ans[iter] = T{"AND", ids[0], ids[1]}
	iter++
	for i := 2; i < N; i++ {
		ans[iter] = T{"AND", iter - 1, ids[i]}
		iter++
	}
	return iter
}
