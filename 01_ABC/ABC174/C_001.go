package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	cnt, pad := 1, 0
	if k%2 == 0 {
		cnt = -1
	} else if k%5 == 0 {
		cnt = -1
	} else {
		digit := 0
		tmp := k
		for tmp > 0 {
			tmp /= 10
			digit++
		}
		pad = digit - 1
		n := 7
		digit--
		for ; digit > 0; digit-- {
			n *= 10
			n += 7
		}
		if n < k {
			n *= 10
			n += 7
			pad++
		}
		for {
			if n%k == 0 {
				break
			}
			n %= k
			n *= 10
			n += 7
			cnt++
		}
	}

	fmt.Println(cnt + pad)
}
