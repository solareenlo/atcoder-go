package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	var a [4]point
	for i := 0; i < 4; i++ {
		fmt.Scan(&a[i].x, &a[i].y)
	}

	for i := 0; i < 4; i++ {
		A, B, C := a[i], a[(i+1)%4], a[(i+2)%4]
		if (B.x-A.x)*(C.y-B.y)-(B.y-A.y)*(C.x-B.x) < 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
