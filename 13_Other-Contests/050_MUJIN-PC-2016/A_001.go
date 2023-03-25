package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c := scanner.Text()[0]
	right := map[byte]bool{'O': true, 'K': true, 'L': true, 'P': true}

	if right[c] {
		fmt.Println("Right")
	} else {
		fmt.Println("Left")
	}
}
