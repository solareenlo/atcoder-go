package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < 39; i++ {
		for j := 1; j < 27; j++ {
			three := 1
			five := 1
			cnt := i
			for k := 0; k < cnt; k++ {
				three *= 3
			}
			cnt = j
			for k := 0; k < cnt; k++ {
				five *= 5
			}
			if three+five == n {
				fmt.Println(i, j)
				return
			}
		}
	}
	fmt.Println(-1)
}
