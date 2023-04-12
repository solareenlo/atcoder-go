package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	ma := 0
	sum := 0
	var a, b int
	fmt.Fscan(in, &a, &b)
	bo := false
	boo := true
	if a == 1 || b == 1 {
		ma2 := 0
		ma3 := 0
		for i := 0; i < max(a, b); i++ {
			var c int
			fmt.Fscan(in, &c)
			if c == 5 {
				ma = max(ma, ma2)
				ma2 = 0
				bo = true
			} else {
				if c != 0 {
					boo = false
				}
				if !bo {
					ma3 = max(ma3, c)
				}
				ma2 = max(ma2, c)
			}
		}
		ma = max(ma, ma2)
		ma2 = min(ma2, ma3)
		for ma2 > 5 {
			ma2 = (ma2 * 2) % 10
			sum++
		}
	} else {
		for i := 0; i < a; i++ {
			for j := 0; j < b; j++ {
				var c int
				fmt.Fscan(in, &c)
				ma = max(ma, c)
				if c == 5 {
					bo = true
				}
			}
		}
	}
	if bo {
		for ma > 5 {
			ma = (ma * 2) % 10
			sum++
		}
		fmt.Println("Yes", sum+1)
	} else if ma == 0 && boo {
		fmt.Println("Yes", 0)
	} else {
		fmt.Println("No")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
