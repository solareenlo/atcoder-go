package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	for i := 0; i <= len(S); i++ {
		for j := i; j <= len(S); j++ {
			s := S[:i] + S[j:]
			if s == "keyence" {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}
