package main

import (
	"bufio"
	"fmt"
	"os"
)

var out = bufio.NewWriter(os.Stdout)

var N int

func main() {
	defer out.Flush()

	fmt.Scan(&N)
	id := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i + 1
	}
	def := make([][]int, 1)
	def[0] = id
	Solve(def)
}

func Solve(par [][]int) {
	if len(par) == N {
		fmt.Fprint(out, "!")
		out.Flush()
		for _, a := range par {
			fmt.Fprintf(out, " %d", a[0])
			out.Flush()
		}
		fmt.Fprintln(out)
		out.Flush()
		os.Exit(0)
	}

	ask(par)

	nex := make([][]int, 0)

	for _, a := range par {
		A := make([]int, 0)
		B := make([]int, 0)
		C := make([]int, 0)
		if len(a) == 1 {
			nex = append(nex, a)
		} else if (len(a) & 1) != 0 {
			for q := 0; q < len(a)-1; q++ {
				rotate(a, 1)
				z := ask(par)
				if z == 1 {
					A = append(A, a[len(a)-1])
				} else if z == 0 {
					B = append(B, a[len(a)-1])
				} else {
					C = append(C, a[len(a)-1])
				}
			}
			if len(A) == len(C) {
				B = append(B, a[0])
			} else if len(A) < len(C) {
				A = append(A, a[0])
			} else {
				C = append(C, a[0])
			}
		} else {
			for q := 0; q < len(a)-2; q++ {
				tmp := a[1:]
				rotate(tmp, 1)
				z := ask(par)
				if z == 1 {
					A = append(A, a[len(a)-1])
				} else if z == 0 {
					B = append(B, a[len(a)-1])
				} else {
					C = append(C, a[len(a)-1])
				}
			}
			if len(A) == len(C) {
				B = append(B, a[1])
			} else if len(A) < len(C) {
				A = append(A, a[1])
			} else {
				C = append(C, a[1])
			}
			rotate(a, 1)
			z := ask(par)
			if z == 1 {
				A = append(A, a[len(a)-1])
			} else {
				C = append(C, a[len(a)-1])
			}
		}
		if len(A) != 0 {
			nex = append(nex, A)
		}
		if len(B) != 0 {
			nex = append(nex, B)
		}
		if len(C) != 0 {
			nex = append(nex, C)
		}
	}

	Solve(nex)
}

func rotate(slice []int, middle int) {
	if middle == 0 || middle == len(slice) {
		return
	}

	reverse(slice[:middle])
	reverse(slice[middle:])
	reverse(slice)
}

func reverse(slice []int) {
	i := 0
	j := len(slice) - 1

	for i < j {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
}

func ask(par [][]int) int {
	fmt.Fprint(out, "?")
	out.Flush()
	for _, a := range par {
		for _, b := range a {
			fmt.Fprintf(out, " %d", b)
			out.Flush()
		}
	}
	fmt.Fprintln(out)
	out.Flush()
	var c string
	fmt.Scan(&c)
	if c == "+" {
		return 1
	} else if c == "-" {
		return -1
	} else {
		return 0
	}
}
