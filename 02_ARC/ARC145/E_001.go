package main

import (
	"bufio"
	"fmt"
	"os"
)

const B = 60

var bs [B]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 1005

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	pa := make([]int, N)
	pb := make([]int, N)
	var id [N][2]int
	tp := 0
	for i := 1; i <= n; i++ {
		if ins(b[i]^a[i], &pb[i]) != 0 {
			fmt.Fprintln(out, "No")
			return
		}
		v := ins(a[i], &pa[i])
		if v != 0 {
			v--
			tp++
			id[tp][0] = i
			id[tp][1] = v
		}
		pb[i] ^= pa[i]
	}

	top := 0
	stk := make([]int, 70005)
	for i := n; i > 1; i-- {
		for j := tp; j > 0; j-- {
			p := id[j][0]
			x := id[j][1]
			v := 0
			for j := 1; j <= i; j++ {
				v ^= pb[j]
			}
			if (((v ^ pa[i]) >> x) & 1) != 0 {
				top++
				stk[top] = p + 1
				for j := 1; j <= p+1; j++ {
					pb[j] ^= pb[j-1]
				}
			}
		}
		top++
		stk[top] = i
		for j := 1; j <= i; j++ {
			pb[j] ^= pb[j-1]
		}
	}

	fmt.Fprintln(out, "Yes")
	fmt.Fprintln(out, top)
	for i := top; i > 0; i-- {
		fmt.Fprintf(out, "%d ", stk[i])
	}
}

func ins(x int, p *int) int {
	*p = 0
	for i := B - 1; i >= 0; i-- {
		if ((x >> i) & 1) != 0 {
			if bs[i] == 0 {
				bs[i] = x
				*p ^= (1 << i)
				return i + 1
			}
			x ^= bs[i]
			*p ^= 1 << i
		}
	}
	return 0
}
