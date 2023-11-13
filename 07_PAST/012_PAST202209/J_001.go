package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = 3.141592653589

	var H, W, D float64
	fmt.Scan(&H, &W, &D)
	if H*H+W*W <= 4*D*D {
		fmt.Printf("1.00000000\n")
		return
	}
	ans := PI * D * D
	if H < 2*D {
		ans -= 2 * math.Acos(H/(2*D)) * D * D
		ans += H * math.Sqrt(D*D-H*H/4)
	}
	if W < 2*D {
		ans -= 2 * math.Acos(W/(2*D)) * D * D
		ans += W * math.Sqrt(D*D-W*W/4)
	}
	ans /= (H * W)
	fmt.Printf("%.8f\n", ans)
}
