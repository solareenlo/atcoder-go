package main

import "fmt"

func main() {
	const N = 200005

	var n, mo int
	fmt.Scan(&n, &mo)
	var a [N]int
	a[0] = 1
	for i := 1; i <= n; i++ {
		a[i] = a[i-1] * n % mo
	}
	su := 1
	ans := 0
	for i := 1; i <= n; i++ {
		su = su * (n - i + 1) % mo
		ans = ans + su*(((i-1)*i/2)%mo)%mo*a[n-i]%mo
	}
	fmt.Println(ans % mo)
}
