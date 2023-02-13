package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)
	var op string
	var num int
	i := 0
	for i < n {
		fmt.Scan(&op, &num)
		if op == "/" && num == 0 {
			fmt.Println("error")
			break
		}
		i++
		switch op {
		case "+":
			a += num
		case "-":
			a -= num
		case "*":
			a *= num
		case "/":
			a /= num
		}
		fmt.Print(i, ":", a, "\n")
	}
}
