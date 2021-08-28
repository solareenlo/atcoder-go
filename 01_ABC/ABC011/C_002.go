package main

import "fmt"

func main() {
	var n, ng1, ng2, ng3 int
	fmt.Scan(&n, &ng1, &ng2, &ng3)
	if n != ng1 && n != ng2 && n != ng3 {
		for i := 0; i < 100; i++ {
			if n-3 != ng1 && n-3 != ng2 && n-3 != ng3 {
				n -= 3
			} else if n-2 != ng1 && n-2 != ng2 && n-2 != ng3 {
				n -= 2
			} else if n-1 != ng1 && n-1 != ng2 && n-1 != ng3 {
				n -= 1
			} else {
				break
			}
		}
	}
	if n <= 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
