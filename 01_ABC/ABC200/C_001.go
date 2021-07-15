package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scan(&n)
	s := make([]int64, 200)
	res := int64(0)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		res += s[a%200]
		s[a%200]++
	}
	fmt.Println(res)
}
