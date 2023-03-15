package main

import "fmt"

func main() {
	const N = 1000000

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > b {
		a, b = b, a
	}

	if b > c {
		fmt.Println("NO")
		return
	}

	minp := make([]int, N+1)
	primes := make([]int, 0)
	for i := 2; i <= N; i++ {
		if minp[i] == 0 {
			minp[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > N {
				break
			}
			minp[i*p] = p
			if minp[i] == p {
				break
			}
		}
	}

	var get func(int, int) int
	get = func(n, p int) int {
		n /= p
		ans := 0
		for n > 0 {
			ans += n
			n /= p
		}
		return ans
	}

	for _, p := range primes {
		if get(a, p)+get(b, p) > get(c, p) {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
