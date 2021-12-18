package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	mp := map[string]bool{}
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		mp[sc.Text()] = true
	}
	fmt.Println(len(mp))
}
