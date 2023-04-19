package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var s, t string
	fmt.Fscan(in, &s, &t)
	s = reverseString(s)
	t = reverseString(t)
	a := new(big.Int)
	a, _ = a.SetString(s, 2)
	b := new(big.Int)
	b, _ = b.SetString(t, 2)
	mask := new(big.Int)
	if countOnes(a) != countOnes(b) {
		fmt.Println("MuriyarokonnNaN")
		return
	}

	for k := 0; k < n; k++ {
		x := countOnes(new(big.Int).And(new(big.Int).Xor(a, b), mask))
		y := countOnes(new(big.Int).And(new(big.Int).Xor(a, b), new(big.Int).Not(mask)))
		if y <= x {
			fmt.Println(k)
			return
		}
		t := a.Bit(n - 1)
		a.Lsh(a, 1)
		a.SetBit(a, 0, t)
		a.SetBit(a, n, 0)
		mask.SetBit(mask, k, 1)
	}
}

func countOnes(n *big.Int) int {
	count := 0
	for _, digit := range n.Text(2) {
		if digit == '1' {
			count++
		}
	}
	return count
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
