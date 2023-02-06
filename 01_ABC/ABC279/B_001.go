package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t string
	fmt.Fscan(in, &s, &t)

	if strings.Index(s, t) > -1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
