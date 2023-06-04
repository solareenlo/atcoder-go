package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	line := in.Text()
	regex := regexp.MustCompile(`(\s)+`)
	ans := regex.ReplaceAllString(line, ",")
	fmt.Println(ans)
}
