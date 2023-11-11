package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 107

	var a, t [N]int

	var n int
	fmt.Fscan(in, &n)
	sum := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	if sum%n != 0 && sum%n != (n*(n+1)/2)%n {
		fmt.Println("No")
		return
	}
	pd := 0
	if sum%n != 0 {
		pd = 1
		sum += n * (n + 1) / 2
		for i := 1; i <= n; i++ {
			a[i] += i
		}
	}
	sum /= n
	ans := 0
	for i := 1; i <= n; i++ {
		if a[i] > sum {
			ans += a[i] - sum
		}
	}
	fmt.Println("Yes")
	fmt.Println(ans + ans + pd)
	if pd != 0 {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
	for ans != 0 {
		x, y, cnt := 0, 0, 2
		for i := 1; i <= n; i++ {
			if a[i] < sum {
				x = i
			}
			if a[i] > sum {
				y = i
			}
		}
		t[x] = 1
		t[y] = 2
		for i := 1; i <= n; i++ {
			if i != x && i != y {
				cnt++
				t[i] = cnt
			}
		}
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", n-t[i]+1)
		}
		fmt.Println()
		t[x] = 2
		t[y] = 1
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", t[i])
		}
		fmt.Println()
		a[x]++
		a[y]--
		ans--
	}
}
