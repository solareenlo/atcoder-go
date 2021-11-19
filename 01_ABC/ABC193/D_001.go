package main

import "fmt"

func main() {
	var n int
	var a, b string
	fmt.Scan(&n, &a, &b)

	t1, t2 := [10]int{}, [10]int{}
	for i := 0; i < 4; i++ {
		t1[a[i]-'0']++
		t2[b[i]-'0']++
	}

	sum1, sum2 := 0, 0
	for i := 0; i < 10; i++ {
		sum1 += i * pow(10, t1[i])
		sum2 += i * pow(10, t2[i])
	}

	res := 0.0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			A := sum1 + 9*i*pow(10, t1[i])
			T := sum2 + 9*j*pow(10, t2[j])
			if A > T {
				if i == j {
					res += 1.0 * float64(n-t1[i]-t2[i]) / float64(9*n-8) * float64(n-t1[j]-t2[j]-1) / float64(9*n-9)
				} else {
					res += 1.0 * float64(n-t1[i]-t2[i]) / float64(9*n-8) * float64(n-t1[j]-t2[j]) / float64(9*n-9)
				}
			}
		}
	}
	fmt.Println(res)
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
