package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	k := 0
	for i := K; i > 0; i /= 10 {
		k = k*10 + i%10
	}

	ans := 0
	if k < K {
		fmt.Println(0)
	} else {
		if k != K {
			for k <= N {
				ans++
				k *= 10
			}
		}
		for K <= N {
			ans++
			K *= 10
		}
		fmt.Println(ans)
	}
}
