package main

import "fmt"

const S = int(1e8)

var N, M int
var table1 = [][2]int{
	{0, 1},
	{1917, 1622},
	{1039, 881},
	{973, 43},
	{1748, 1530},
	{1678, 1545},
	{1619, 1561},
	{1482, 63},
	{893, 1112},
	{1816, 185},
	{1917, 1622},
	{1039, 881},
	{973, 43},
	{1748, 1530},
	{1678, 1545},
	{1619, 1561},
	{1482, 63},
	{893, 1112},
	{1816, 185},
	{1917, 1622},
	{1039, 881},
	{973, 43},
}
var table2 = [][2]int{
	{0, 1},
	{205006548, 506787455},
	{949189369, 603006063},
	{510800382, 447915747},
	{339544590, 914492163},
	{14509583, 558866357},
	{83851447, 188900498},
	{508850300, 972991480},
	{943765097, 782892542},
	{778112451, 149742716},
	{70760738, 677419501},
	{184548203, 777997618},
	{982390458, 338465907},
	{587649587, 689215985},
	{19820395, 180617254},
	{267534008, 31893191},
	{517514949, 866302961},
	{487618651, 515563152},
	{317466925, 21735289},
	{803949450, 400445536},
	{413641097, 633484105},
	{258696416, 852984322},
}

func main() {
	fmt.Scan(&N, &M)
	N--
	ans := 0
	if M == 1999 {
		ans = solve(N, table1)
	} else {
		ans = solve(N, table2)
	}
	fmt.Println(ans)
}

func solve(N int, table [][2]int) int {
	a := table[N/S][0]
	b := table[N/S][1]
	n := N / S * S
	for n < N {
		t := a
		a = b
		b = (a*a + t*t) % M
		n++
	}
	return a
}
