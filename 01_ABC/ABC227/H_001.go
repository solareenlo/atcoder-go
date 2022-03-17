package main

import (
	"bufio"
	"fmt"
	"os"
)

func F(n int, s string) {
	for i := 0; i < n; i++ {
		fmt.Print(s)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	f := make([]int, 10)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Fscan(in, &f[i*3-3+j])
			if i != 2 || j != 2 {
				f[i*3-3+j]--
			}
		}
	}

	for i := 0; i < f[1]+1; i++ {
		for j := 0; j < f[7]+1; j++ {
			for k := 0; k < f[3]+1; k++ {
				for l := 0; l < f[9]+1; l++ {
					L := f[4] - i - j
					R := f[6] - k - l
					U := f[2] - (f[1] - i + f[3] - k)
					D := f[8] - (f[7] - j + f[9] - l)
					if R >= 0 && L >= 0 && U >= 0 && D >= 0 && R+L+U+D == f[5] {
						F(1, "D")
						F(i, "UD")
						F(L, "RL")
						F(1, "D")
						F(j, "UD")
						F(1, "R")
						F(f[7]-j, "LR")
						F(D, "UD")
						F(1, "R")
						F(f[9]-l, "LR")
						F(1, "U")
						F(l, "DU")
						F(R, "LR")
						F(1, "U")
						F(k, "DU")
						F(f[3]-k, "LR")
						F(1, "L")
						F(U, "DU")
						F(f[1]-i, "LR")
						F(1, "L\n")
						return
					}
				}
			}
		}
	}
	fmt.Println("NO")
}
