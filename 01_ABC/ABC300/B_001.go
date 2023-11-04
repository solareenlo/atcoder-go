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
	A := make([]string, h)
	B := make([]string, h)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	for i := range B {
		fmt.Fscan(in, &B[i])
	}
	s := 0
	p := false
	for ; s < h; s++ {
		for t := 0; t < w; t++ {
			f := true
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if f && A[(i+h-s)%h][(j+w-t)%w] == B[i][j] {
						f = true
					} else {
						f = false
					}
				}
			}
			if !p && !f {
				p = false
			} else {
				p = true
			}
		}
	}
	if p {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
