package main

import "fmt"

func calc() {
	fmt.Println("rem D C")
	fmt.Println("mul D D 4915446")
	fmt.Println("rem D D")
	fmt.Println("mul D D 1000000007")
	fmt.Println("add D D C")
	fmt.Println("mul C D 996491785301655553")
}

func main() {
	fmt.Println("26")
	fmt.Println("mul C A B")
	calc()
	fmt.Println("mul C C 343639189")
	calc()
	calc()
	calc()
}
