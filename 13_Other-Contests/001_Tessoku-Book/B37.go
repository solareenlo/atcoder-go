package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var s [10]int
	for i := 1; i < 10; i++ {
		s[i] = s[i-1] + i
	}
	i := 10
	sum := 45*(N/10) + s[N%10]
	for N > i {
		p := (N / i) % 10
		if p == 0 {
			sum += 45*(N/(i*10))*i + 0*i + p*(N%i+1)
		} else {
			sum += 45*(N/(i*10))*i + s[p-1]*i + p*(N%i+1)
		}
		i *= 10
	}
	fmt.Println(sum)
}
