package main

import (
	"fmt"
	"math/big"
)

func main() {
	N := new(big.Int)
	fmt.Scan(N)
	mod := new(big.Int)
	if mod.Mod(N, big.NewInt(6)).Cmp(big.NewInt(0)) != 0 {
		fmt.Println("nO")
	} else {
		fmt.Println("yES")
	}
	if mod.Mod(N, big.NewInt(11)).Cmp(big.NewInt(0)) != 0 {
		fmt.Println("nO")
	} else {
		fmt.Println("yES")
	}
}
