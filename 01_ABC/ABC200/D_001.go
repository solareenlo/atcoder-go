package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	b := make([][]int, 200)
	for i := range b {
		b[i] = make([]int, 0)
	}

	cnt := min(n, 8)
	for bit := 1; bit < 1<<cnt; bit++ {
		sig := 0
		s := make([]int, 0)
		for i := 0; i < cnt; i++ {
			if bit&(1<<i) > 0 {
				s = append(s, i+1)
				sig += a[i]
				sig %= 200
			}
		}
		if len(b[sig]) != 0 {
			fmt.Println("Yes")
			output(b[sig])
			output(s)
			return
		} else {
			b[sig] = s
		}
	}
	fmt.Println("No")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func output(a []int) {
	fmt.Print(len(a))
	for _, x := range a {
		fmt.Print(" ", x)
	}
	fmt.Println()
}
