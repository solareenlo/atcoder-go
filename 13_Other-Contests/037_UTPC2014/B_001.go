package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	fmt.Scan(&x, &y)

	fmt.Printf("%.0f %.0f %.0f %.0f\n", math.Floor(x), math.Floor(y+1), math.Floor(x)+math.Round(1000*(x-math.Floor(x))), math.Floor(y+1)+math.Round(1000*(y-math.Floor(y+1))))
	fmt.Printf("%.0f %.0f %.0f %.0f\n", math.Floor(x-1), math.Floor(y), math.Floor(x-1)+math.Round(1000*(x-math.Floor(x-1))), math.Floor(y)+math.Round(1000*(y-math.Floor(y))))
}
