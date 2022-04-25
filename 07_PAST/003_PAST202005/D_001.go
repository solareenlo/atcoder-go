package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(a []string) []int {
	res := make([]int, 0)
	for i := 1; i < len(a[0]); i += 4 {
		t := 0
		for _, s := range a {
			for _, c := range s[i : i+4] {
				t = t*2 + int(c)%2
			}
		}
		res = append(res, t%23)
	}
	return res
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(IN, &n)
	in := make([]string, 5)
	for i := range in {
		fmt.Fscan(IN, &in[i])
	}
	for _, x := range f(in) {
		fmt.Fprint(out, string("5 4 90   7 316 2 8"[x]))
	}
}
