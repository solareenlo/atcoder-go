package main

import "fmt"

func main() {
	type pair struct {
		x, y int
	}

	var N int
	fmt.Scan(&N)
	kotae := make([]pair, 0)
	chousei := 7
	var a, b int
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j += chousei {
			fmt.Printf("? ")
			for k := 0; k < N; k++ {
				if k == i || (j <= k && k < j+chousei) {
					a = 1
				} else {
					a = 0
				}
				fmt.Printf("%d", a)
			}
			fmt.Printf("\n")
			fmt.Scan(&a)
			if a != 0 {
				fmt.Printf("? ")
				for k := 0; k < N; k++ {
					if j <= k && k < j+chousei {
						b = 1
					} else {
						b = 0
					}
					fmt.Printf("%d", b)
				}
				fmt.Printf("\n")
				fmt.Scan(&b)
				if a == b {
					continue
				}
				for k := j; k < min(j+chousei, N); k++ {
					fmt.Printf("? ")
					for l := 0; l < N; l++ {
						if l == i || k == l {
							b = 1
						} else {
							b = 0
						}
						fmt.Printf("%d", b)
					}
					fmt.Printf("\n")
					fmt.Scan(&b)
					if b != 0 {
						kotae = append(kotae, pair{i, k})
					}
				}
			}
		}
	}
	fmt.Printf("!")
	for i := 0; i < len(kotae); i++ {
		fmt.Printf(" (%d,%d)", kotae[i].x, kotae[i].y)
	}
	fmt.Printf("\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
