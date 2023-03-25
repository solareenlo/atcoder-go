package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for T > 0 {
		var a int
		fmt.Scan(&a)
		if a == 0 || a == 1 {
			fmt.Println("No")
			T--
			continue
		}
		if a == 3 {
			fmt.Println("5")
			fmt.Println("3 3 4 4 6")
		} else if a == 2 {
			fmt.Println("5")
			fmt.Println("0 1 1 1 2")
		} else if a == 4 {
			fmt.Println("5")
			fmt.Println("0 1 1 2 4")
		} else if a == 8 {
			fmt.Println("5")
			fmt.Println("0 1 1 4 8")
		} else if a%4 == 0 {
			fmt.Printf("5\n%d 1 0 %d %d\n", 4, (a-1)/4, a-4)
		} else {
			fmt.Printf("5\n%d 1 0 %d %d\n", a%4, a/4, a-a%4)
		}
		T--
	}
}
