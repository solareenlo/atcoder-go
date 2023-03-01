package main

import (
	"bufio"
	"fmt"
	"os"
)

var dong [30]int = [30]int{0, 1, 2, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
var UTPC [4]int = [4]int{0, 0, 1, 0}

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	for i := 0; i < 4; i++ {
		if dong[s[i]-'A'+1] != UTPC[i] {
			fmt.Println("no")
			return
		}
	}
	fmt.Println("yes")
}
