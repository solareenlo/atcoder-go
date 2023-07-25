package main

import "fmt"

func main() {
	x := 0
	y := 0
	now := 2
	for now <= 200 {
		if y+now > 1500 {
			x += now - 2
			y = 0
		}
		fmt.Println(x+now/2, y+now/2)
		y += now
		now += 2
	}
}
