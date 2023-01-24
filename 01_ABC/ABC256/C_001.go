package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	if (a+b+c)^(d+e+f) != 0 {
		fmt.Println(0)
		os.Exit(0)
	}

	res := 0
	for i := 1; i <= 30; i++ {
		for j := 1; j <= 30; j++ {
			if i+j < a {
				for k := 1; k <= 30; k++ {
					if i+k < d {
						for l := 1; l <= 30; l++ {
							if k+l < b {
								if j+l < e {
									if a-i-j+b-k-l < f {
										res++
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(res)
}
