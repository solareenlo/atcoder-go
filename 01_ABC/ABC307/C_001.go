package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f func() int
	f = func() int {
		ret := 0
		var s string
		fmt.Fscan(in, &s)
		for _, c := range s {
			if c == '#' {
				ret = (ret << 1) | 1
			} else {
				ret = (ret << 1) | 0
			}
		}
		return ret
	}
	var a, b [15]int
	var ha, wa int
	fmt.Fscan(in, &ha, &wa)
	for i := 0; i < ha; i++ {
		a[i] = f()
	}
	var hb, wb int
	fmt.Fscan(in, &hb, &wb)
	for i := 0; i < hb; i++ {
		b[i] = f()
	}
	var hx, wx int
	fmt.Fscan(in, &hx, &wx)
	c := make([]int, 30)
	for i := 0; i < hx; i++ {
		c[i+10] = (f() << 10)
	}
	for i := 0; i < 21; i++ {
		for j := 0; j < 21; j++ {
			for k := 0; k < 21; k++ {
				for l := 0; l <= 20; l++ {
					m := make([]int, 30)
					for x := 0; x < 10; x++ {
						m[x+i] |= a[x] << j
						m[x+k] |= b[x] << l
					}
					if reflect.DeepEqual(m, c) {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}
	fmt.Println("No")
}
