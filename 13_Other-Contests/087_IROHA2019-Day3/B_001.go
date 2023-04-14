package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var d [2][2][3]int
	fmt.Fprintln(out, "? 0 1 2")
	out.Flush()
	var s, t string
	fmt.Scan(&s, &t)
	d[0][0][f(s)]++
	d[0][1][f(s)]++
	d[1][0][f(t)]++
	d[1][1][f(t)]++
	for i := 3; i <= 7; i++ {
		fmt.Fprintln(out, "? 0 1", i)
		out.Flush()
		fmt.Scan(&s, &t)
		d[0][0][f(s)]++
		d[1][0][f(t)]++
	}
	for i := 3; i <= 7; i++ {
		fmt.Fprintln(out, "? 0 2", i)
		out.Flush()
		fmt.Scan(&s, &t)
		d[0][1][f(s)]++
		d[1][1][f(t)]++
	}
	b := [2]int{0, 0}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			s := d[i][j][0]
			r := d[i][j][1]
			if s == 4 {
				if r == 2 {
					b[i]++
				}
			}
			if s == 2 {
				if r == 2 {
					b[i]++
				}
			}
			if s == 0 {
				if r == 6 {
					b[i]++
				}
			}
		}
	}
	if b[0] == 2 {
		fmt.Fprintln(out, "! 1")
		out.Flush()
	} else {
		fmt.Fprintln(out, "! 2")
		out.Flush()
	}
}

func f(s string) int {
	if s[0] == 'S' {
		return 0
	}
	if s[0] == 'R' {
		return 1
	}
	return 2
}
