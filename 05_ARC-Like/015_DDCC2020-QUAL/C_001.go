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

	var h, w, k int
	fmt.Fscan(in, &h, &w, &k)

	B := make([]string, h)
	for i := range B {
		fmt.Fscan(in, &B[i])
	}

	ans := [300][300]int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if B[i][j] == '#' {
				ans[i][j] = k
				k--
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w-1; j++ {
			if ans[i][j+1] == 0 {
				ans[i][j+1] = ans[i][j]
			}
		}
		for j := w - 1; j > 0; j-- {
			if ans[i][j-1] == 0 {
				ans[i][j-1] = ans[i][j]
			}
		}
	}

	for j := 0; j < w; j++ {
		for i := 0; i < h-1; i++ {
			if ans[i+1][j] == 0 {
				ans[i+1][j] = ans[i][j]
			}
		}
		for i := h - 1; i > 0; i-- {
			if ans[i-1][j] == 0 {
				ans[i-1][j] = ans[i][j]
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprint(out, ans[i][j])
			if j < w-1 {
				fmt.Fprint(out, " ")
			} else {
				fmt.Fprint(out, "\n")
			}
		}
	}
}
