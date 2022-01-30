package main

import (
	"fmt"
	"math/big"
)

func main() {
	var s string
	fmt.Scan(&s)

	a := new(big.Int)
	a.SetString(s, 10)

	b := Sub(Mul(Add(a, big.NewInt(1)), Add(a, big.NewInt(1))), big.NewInt(1))
	a = Mul(a, a)
	for Div((Add(a, big.NewInt(99))), big.NewInt(100)).Cmp(Div(b, big.NewInt(100))) != 1 {
		a = Div(Add(a, big.NewInt(99)), big.NewInt(100))
		b = Div(b, big.NewInt(100))
	}
	fmt.Println(a)
}

func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func Div(x, y *big.Int) *big.Int {
	return big.NewInt(0).Div(x, y)
}
func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}
