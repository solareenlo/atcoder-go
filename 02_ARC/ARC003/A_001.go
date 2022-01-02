package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	gpa := map[byte]int{}
	for i := 0; i < n; i++ {
		gpa[s[i]]++
	}

	sum := 0.0
	for k, v := range gpa {
		switch k {
		case 'A':
			sum += 4.0 * float64(v)
		case 'B':
			sum += 3.0 * float64(v)
		case 'C':
			sum += 2.0 * float64(v)
		case 'D':
			sum += 1.0 * float64(v)
		}
	}

	fmt.Println(sum / float64(n))
}
