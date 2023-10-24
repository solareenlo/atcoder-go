package main

import "fmt"

type pair struct {
	x, y int
}

const MOD = 1000000007
const prime_max = 1100000

var NP int
var prime []int
var W, H int
var divp [prime_max]int
var D [80808][]pair
var CX [80808]int
var BS [1200000]int
var A, B [404040]int

func cprime() {
	if NP != 0 {
		return
	}
	for i := 2; i < prime_max; i++ {
		if divp[i] == 0 {
			prime = append(prime, i)
			NP++
			for j := i * i; j >= i && j < prime_max; j += i {
				if divp[j] == 0 {
					divp[j] = i
				}
			}
		}
	}
}

func gendivs_sub(ret *[]pair, D []pair, cur, mul, tot int) {
	if cur == len(D) {
		*ret = append(*ret, pair{mul, tot})
		return
	}
	gendivs_sub(ret, D, cur+1, mul, tot)
	tot *= D[cur].x - 1
	for i := 1; i <= D[cur].y; i++ {
		mul *= D[cur].x
		gendivs_sub(ret, D, cur+1, mul, tot)
		tot *= D[cur].x
	}
}

func gendivs(D []pair) []pair {
	ret := make([]pair, 0)
	gendivs_sub(&ret, D, 0, 1, 1)
	return ret
}

func hoge(X, Y int) int {
	X -= W - 1
	Y -= H - 1

	for x := 0; x < 2*W-1; x++ {
		CX[x] = X + x
		D[x] = make([]pair, 0)
	}
	for _, p := range prime {
		for cur := (X+p-1)/p*p - X; cur < 2*W-1; cur += p {
			D[cur] = append(D[cur], pair{p, 0})
			for CX[cur]%p == 0 {
				CX[cur] /= p
				D[cur][len(D[cur])-1].y++
			}
		}
	}

	for x := 1; x <= prime_max; x++ {
		BS[x] = 0
		for cur := (Y+x-1)/x*x - Y; cur < 2*H-1; cur += x {
			BS[x] += B[cur]
		}
		BS[x] %= MOD
	}
	for x := 0; x < 2*W-1; x++ {
		if CX[x] > 1 {
			D[x] = append(D[x], pair{CX[x], 1})
		}
	}

	ret := 0
	for x := 0; x < 2*W-1; x++ {
		divs := gendivs(D[x])
		for _, d := range divs {
			T := 0
			if d.x <= prime_max {
				T = BS[d.x]
			} else {
				for cur := (Y+d.x-1)/d.x*d.x - Y; cur < 2*H-1; cur += d.x {
					T += B[cur]
				}
			}
			ret += (A[x] * T % MOD) * d.y % MOD
		}
	}
	return ret % MOD
}

func main() {
	cprime()
	var X, Y [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&X[i], &Y[i])
	}
	X[1] -= X[0]
	X[2] -= X[0]
	Y[1] -= Y[0]
	Y[2] -= Y[0]

	fmt.Scan(&W, &H)
	for i := 0; i < 2*W-1; i++ {
		if i <= W-1 {
			A[i] = i + 1
		} else {
			A[i] = 2*W - 1 - i
		}
	}
	for i := 0; i < 2*H-1; i++ {
		if i <= H-1 {
			B[i] = i + 1
		} else {
			B[i] = 2*H - 1 - i
		}
	}

	S := ((X[1]%MOD)*(Y[2]%MOD)%MOD - (X[2]%MOD)*(Y[1]%MOD)%MOD + MOD) % MOD * (MOD + 1) / 2 % MOD
	ret := ((((((S + 1) * W % MOD) * W % MOD) * W % MOD) * H % MOD) * H % MOD) * H % MOD
	tmp := hoge(X[1], Y[1])
	ret += MOD - (W*H*tmp%MOD)*(MOD+1)/2%MOD
	ret += MOD - (W*H*hoge(X[2], Y[2])%MOD)*(MOD+1)/2%MOD
	ret += MOD - (W*H*hoge(X[1]-X[2], Y[2]-Y[1])%MOD)*(MOD+1)/2%MOD
	fmt.Println(ret % MOD)
}
