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
	mp := make(map[int]int)
	for n > 0 {
		n--
		var i int
		fmt.Fscan(in, &i)
		mp[i]++
	}
	ans := 0
	for _, v := range mp {
		ans += v / 2
	}
	fmt.Println(ans)
}
