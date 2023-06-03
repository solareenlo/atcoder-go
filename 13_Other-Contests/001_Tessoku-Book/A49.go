package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)
	p := make([]int, t)
	q := make([]int, t)
	r := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &p[i], &q[i], &r[i])
	}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < t; i++ {
		if rand.Int()%2 != 0 {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}
	}
}
