package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)
	var num_i, num_s, num_o int
	fmt.Fscan(in, &num_i, &num_s, &num_o)
	board := make([]string, 101)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &board[i])
	}

	r0, r1, c0, c1, s0, s1 := 0, 0, 0, 0, 0, 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == 'o' {
				if i%2 == 0 {
					r0++
				} else {
					r1++
				}
				if j%2 == 0 {
					c0++
				} else {
					c1++
				}
				if (i+j)%2 == 0 {
					s0++
				} else {
					s1++
				}
			}
		}
	}
	if s0 != s1 {
		fmt.Println("T")
		return
	}
	if r0%4 != r1%4 || c0%4 != c1%4 {
		fmt.Println("L")
		return
	}

	hoge := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == 'o' {
				if (i+j)%4 == 0 || (i+j)%4 == 1 {
					hoge++
				}
			}
		}
	}
	if hoge%2 != num_o%2 {
		fmt.Println("O")
		return
	}
	piyo := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == 'o' {
				if i%2 == 0 && j%2 == 0 {
					piyo++
				}
			}
		}
	}
	if piyo%2 != (num_o+num_s)%2 {
		fmt.Println("S")
		return
	}
	fmt.Println("I")
}
