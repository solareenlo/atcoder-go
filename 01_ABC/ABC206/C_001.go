package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	cnt := make(map[int]int)
	res := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		res += i - cnt[a]
		cnt[a]++
	}
	fmt.Println(res)
}
