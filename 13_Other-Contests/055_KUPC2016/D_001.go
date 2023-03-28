package main

import "fmt"

func main() {
	var s, t string
	var N int
	fmt.Scan(&N)
	head := true
	for {
		now := false
		for r := 0; r < 4; r++ {
			if head {
				if (r & 1) != 0 {
					fmt.Printf("#%s\n", s)
				} else {
					fmt.Printf(".%s\n", s)
				}
				if (r & 2) != 0 {
					fmt.Printf("#%s\n", t)
				} else {
					fmt.Printf(".%s\n", t)
				}
			} else {
				if (r & 1) != 0 {
					fmt.Printf("%s#\n", s)
				} else {
					fmt.Printf("%s.\n", s)
				}
				if (r & 2) != 0 {
					fmt.Printf("%s#\n", t)
				} else {
					fmt.Printf("%s.\n", t)
				}
			}
			var ans string
			fmt.Scan(&ans)
			if ans == "end" {
				return
			}
			if ans == "T" {
				now = true
				if head {
					if (r & 1) != 0 {
						s = "#" + s
					} else {
						s = "." + s
					}
					if (r & 2) != 0 {
						t = "#" + t
					} else {
						t = "." + t
					}
				} else {
					if (r & 1) != 0 {
						s += "#"
					} else {
						s += "."
					}
					if (r & 2) != 0 {
						t += "#"
					} else {
						t += "."
					}
				}
				break
			}
		}
		if !now {
			head = !head
		}
	}
}
