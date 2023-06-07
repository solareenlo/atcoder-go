package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b [100001]int

	var q, l int
	fmt.Fscan(in, &q, &l)
	Len := 0
	cnt := 0
	for q > 0 {
		q--
		var op string
		fmt.Fscan(in, &op)
		if op == "Push" {
			var n, m int
			fmt.Fscan(in, &n, &m)
			if l-Len < n {
				fmt.Println("FULL")
				return
			}
			cnt++
			a[cnt] = n
			b[cnt] = m
			Len += n
		}
		if op == "Pop" {
			var n int
			fmt.Fscan(in, &n)
			if Len < n {
				fmt.Println("EMPTY")
				return
			}
			Len -= n
			for n > 0 {
				if a[cnt] <= n {
					n -= a[cnt]
					a[cnt] = 0
					b[cnt] = 0
					cnt--
				} else {
					a[cnt] -= n
					n = 0
				}
			}
		}
		if op == "Top" {
			if Len == 0 {
				fmt.Println("EMPTY")
				return
			}
			fmt.Printf("%d\n", b[cnt])
		}
		if op == "Size" {
			fmt.Printf("%d\n", Len)
		}
	}
	fmt.Println("SAFE")
}
