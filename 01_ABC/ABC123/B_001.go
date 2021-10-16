package main

import "fmt"

func main() {
	sum, rem, tmp := 0, 10, 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&tmp)
		sum += (tmp + 9) / 10
		if tmp%10 != 0 && rem > tmp%10 {
			rem = tmp % 10
		}
	}

	fmt.Println(sum*10 - 10 + rem)
}
