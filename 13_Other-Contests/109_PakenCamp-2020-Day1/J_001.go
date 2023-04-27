package main

import "fmt"

func main() {
	for i := 1; i <= 100000; i++ {
		fmt.Println(i*2+1, i*i*2+i*2, i*i*2+i*2+1)
	}
}
