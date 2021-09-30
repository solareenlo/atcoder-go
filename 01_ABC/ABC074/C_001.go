package main

import "fmt"

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	resWS, resS := 100*a, 0
	for i := 0; i < f/(100*a)+1; i++ {
		for j := 0; j < f/(100*b)+1; j++ {
			w := 100*a*i + 100*b*j
			if w == 0 {
				continue
			}
			for n := 0; n < f/c+1; n++ {
				for m := 0; m < f/d+1; m++ {
					s := c*n + d*m
					if w+s > f {
						continue
					}
					if 100*s > e*w {
						continue
					}
					if s*resWS <= resS*(w+s) {
						continue
					}
					resWS = w + s
					resS = s
				}
			}
		}
	}
	fmt.Println(resWS, resS)
}
