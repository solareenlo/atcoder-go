package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	for i := 0; i < N; i++ {
		var H, W int
		fmt.Fscan(in, &H, &W)
		var ans [4][]string
		ans[0] = []string{"A", "B", "C", "A"}
		ans[1] = []string{"B", "C", "A", "B"}
		ans[2] = []string{"C", "A", "B", "C"}
		ans[3] = []string{"A", "B", "C", "A"}
		for y := 0; y < 4; y++ {
			for x := 0; x < W-4; x++ {
				ans[y] = append(ans[y], ans[y][3])
			}
		}
		for y := 0; y < 4; y++ {
			fmt.Fprintln(out, strings.Join(ans[y], ""))
		}
		for y := 0; y < H-4; y++ {
			fmt.Fprintln(out, strings.Join(ans[3], ""))
		}
	}
}
