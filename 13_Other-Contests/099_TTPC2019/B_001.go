package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		var S string
		fmt.Fscan(in, &S)
		reg := regexp.MustCompile(`^.*okyo.*ech.*`)
		if reg.MatchString(S) {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
