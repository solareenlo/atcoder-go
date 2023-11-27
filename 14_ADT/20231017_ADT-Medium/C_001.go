package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	slice := strings.Split(s, "")
	sort.Strings(slice)
	fmt.Println(strings.Join(slice, ""))
}
