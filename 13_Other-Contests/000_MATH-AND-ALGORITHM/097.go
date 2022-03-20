package main

import "fmt"

func main() {
	var L, R int
	fmt.Scan(&L, &R)

	isp := [1 << 20]bool{}
	ispQ := [5 << 17]bool{}
	for i := 2; i < 1<<20; i++ {
		if !isp[i] {
			for j := i + i; j < 1<<20; j += i {
				isp[j] = true
			}
			j := (L + i - 1) / i * i
			if j == i {
				j += i
			}
			for j <= R {
				ispQ[j-L] = true
				j += i
			}
		}
	}

	cnt := 0
	for j := L; j <= R; j++ {
		if j > 1 && !ispQ[j-L] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
