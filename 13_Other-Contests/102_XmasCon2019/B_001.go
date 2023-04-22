package main

import "fmt"

func main() {
	var N int
	var op string
	fmt.Scan(&N, &op)
	if op == "<" {
		if N < 3 {
			fmt.Println("Merry Christmas!")
			return
		}
		fmt.Print(0)
		for i := 0; i < N-1; i++ {
			fmt.Printf(" %d", i+2)
		}
		fmt.Println()
		return
	}
	if op == "=" {
		for i := 0; i < N; i++ {
			fmt.Print(i)
			if i < N-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		return
	}
	if N < 8 {
		fmt.Println("Merry Christmas!")
		return
	}
	fmt.Print("0 2 3 4")
	for i := 0; i < N-6; i++ {
		fmt.Printf(" %d", i*4+7)
	}
	fmt.Printf(" %d %d\n", (N-7)*4+7+1, (N-7)*4+7+3)
}
