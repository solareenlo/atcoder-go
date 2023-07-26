package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var plist [13]int
	plistlen := 0
	for p := 2; 1 < n; p++ {
		if p*p > n {
			p = n
		}
		if n%p == 0 {
			cnt := 0
			for n%p == 0 {
				n /= p
				cnt++
			}
			plist[plistlen] = cnt
			plistlen++
		}
	}
	numAll := 1
	for i := 0; i < plistlen; i++ {
		numAll *= calcAll(plist[i])
	}
	numTypeA := 1
	for i := 0; i < plistlen; i++ {
		numTypeA *= calcTypeA(plist[i])
	}
	numTypeB := 1
	for i := 0; i < plistlen; i++ {
		numTypeB *= calcTypeB(plist[i])
	}
	numTypeC := 1
	for i := 0; i < plistlen; i++ {
		numTypeC *= calcTypeC(plist[i])
	}
	numPIE := numAll - 2*numTypeA - 2*numTypeB - 4*numTypeC + 7
	fmt.Println(numPIE / 8)
}

func calcAll(n int) int   { return 2*n*n + 2*n + 1 }
func calcTypeA(n int) int { return 2*n + 1 }
func calcTypeB(n int) int { return 2*(n/2) + 1 }
func calcTypeC(n int) int { return 2*(n/3) + 1 }
