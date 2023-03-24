package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := 0
	fmt.Sscan(scanner.Text(), &n)
	for i := 0; i < n && scanner.Scan(); i++ {
		s := scanner.Text()
		ans := 0
		for i := 0; i < len(s); i++ {
			if i+4 < len(s) && s[i:i+5] == "kyoto" {
				ans++
				i += 4
			}
			if i+4 < len(s) && s[i:i+5] == "tokyo" {
				ans++
				i += 4
			}
		}
		fmt.Println(ans)
	}
}
