package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cntP, cntM, cntO := 0, 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1':
			cntO++
		case '+':
			cntP++
		case '-':
			cntM++
		}
	}
	fmt.Println(1 + cntP - cntM)
}
