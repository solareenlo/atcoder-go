package main

import "fmt"

func main() {
	var a [10086]int
	var n int
	fmt.Scan(&n)
	num := 0
	for n > 0 {
		num++
		a[num] = n % 2
		n /= 2
	}
	for i := 1; i <= num; i++ {
		if a[i] != a[num-i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
