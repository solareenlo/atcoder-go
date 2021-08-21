package main

import "fmt"

var d [17]string = [17]string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW", "N"}
var w [13]float64 = [13]float64{0.25, 1.55, 3.35, 5.45, 7.95, 10.75, 13.85, 17.15, 20.75, 24.45, 28.45, 32.65, 12000}

func main() {
	var deg int
	var dis float64
	fmt.Scan(&deg, &dis)

	var i, j int
	for i = 0; i < 17; i++ {
		if deg*10 < 1125+2250*i {
			break
		}
	}
	for j = 0; j < 13; j++ {
		if dis < w[j]*60 {
			break
		}
	}

	if j != 0 {
		fmt.Print(d[i])
	} else {
		fmt.Print("C")
	}
	fmt.Println("", j)
}
