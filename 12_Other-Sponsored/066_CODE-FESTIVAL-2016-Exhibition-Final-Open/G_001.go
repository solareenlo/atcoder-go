package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)

	var comb [606]int
	comb[0] = 1
	for i := 1; i < 600; i++ {
		comb[i] = comb[i-1] * (i + 7) / i
	}

	var cnt [606]int
	for i := 600 - 1; i >= 0; i-- {
		cnt[i] = K / comb[i]
		K %= comb[i]
	}

	for i := 0; i < 600; i++ {
		fmt.Print("FESTIVA")
		for j := 0; j < cnt[i]; j++ {
			fmt.Print("L")
		}
	}
	fmt.Println()
}
