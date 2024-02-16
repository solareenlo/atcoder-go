package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	data := make([][2]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &data[i][0], &data[i][1])
	}
	var M int
	fmt.Fscan(in, &M)
	S := ""
	for i := 0; i < M; i++ {
		var n string
		fmt.Fscan(in, &n)
		for i := 0; i < N; i++ {
			if n == data[i][0] {
				n = data[i][1]
				break
			}
		}
		S += n
	}
	fmt.Println(S)
}
