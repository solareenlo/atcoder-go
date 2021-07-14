package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	res := 0
	for i := 0; i < 10000; i++ {
		flag := make([]bool, 10)
		x := i
		for j := 0; j < 4; j++ {
			flag[x%10] = true
			x /= 10
		}
		flag2 := true
		for j := range s {
			if (s[j] == 'o' && !flag[j]) || (s[j] == 'x' && flag[j]) {
				flag2 = false
			}
		}
		if flag2 {
			res++
		}
	}
	fmt.Println(res)
}
