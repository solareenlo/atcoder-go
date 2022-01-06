package main

import "fmt"

func main() {
	var m, n, N int
	fmt.Scan(&m, &n, &N)

	sale := N
	notSale := 0
	for (sale+notSale)/m > 0 {
		div := (sale + notSale) / m
		rem := (sale + notSale) % m
		notSale = 0
		sale = div * n
		notSale = rem
		N += sale
	}

	fmt.Println(N)
}
