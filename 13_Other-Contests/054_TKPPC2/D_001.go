package main

import (
	"fmt"
	"math/big"
)

func main() {
	var h, w, k int64
	fmt.Scan(&h, &w, &k)
	f := new(big.Int).GCD(nil, nil, big.NewInt(w), big.NewInt(k))
	k /= f.Int64()
	w /= f.Int64()
	fmt.Println((w - 1) * (k - 1) / 2)
}
