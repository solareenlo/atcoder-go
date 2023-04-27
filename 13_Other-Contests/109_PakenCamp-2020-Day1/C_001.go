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

	cnt := 0
	people := make(map[string]int)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for j := 1; j <= x; j++ {
			var tmp string
			fmt.Fscan(in, &tmp)
			people[tmp]++
			if people[tmp] == n {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
