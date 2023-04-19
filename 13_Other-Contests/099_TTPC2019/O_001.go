package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)
	var ans [32][]string
	for i := range ans {
		ans[i] = make([]string, 32)
		for j := range ans[i] {
			ans[i][j] = "#"
		}
	}
	ans[1][1] = "S"
	ans[1][30] = "G"
	for i := 2; i <= 30; i++ {
		ans[i][1] = "."
		ans[i][30] = "."
	}

	i := 1
	for e := 11; e >= 0; e-- {
		if n&(1<<e) != 0 {
			for j := 2; j < 30; j++ {
				ans[i][j] = "."
			}
			for j := 0; j <= e; j++ {
				ans[i+1][3+j] = "."
			}
			i += 3
		}
	}

	fmt.Println(32, 32)
	for _, s := range ans {
		fmt.Fprintln(out, strings.Join(s, ""))
	}
}
