package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	flip := H > W
	if flip {
		H, W = W, H
	}

	if H > 3 || (H == 3 && W > 3) {
		fmt.Println("No")
		return
	} else if H == 3 {
		fmt.Println("Yes")
		fmt.Println("255")
		fmt.Println("252")
		fmt.Println("552")
		return
	} else if H == 2 {
		if W%7 > 1 && W%7 < 6 {
			fmt.Println("No")
			return
		} else {
			s1, s2 := "", ""
			r1, r2 := "552555", "555255"
			s1, s2 = generateRows(W, r1, r2)
			fmt.Println("Yes")
			if flip {
				for i := 0; i < W; i++ {
					fmt.Printf("%c%c\n", s1[i], s2[i])
				}
			} else {
				fmt.Println(s1)
				fmt.Println(s2)
			}
		}
	} else {
		if W%7 != 0 && W%7 != 2 && W%7 != 5 {
			fmt.Println("No")
			return
		} else {
			s := generateRow(W)
			fmt.Println("Yes")
			if flip {
				for _, c := range s {
					fmt.Printf("%c\n", c)
				}
			} else {
				fmt.Println(s)
			}
		}
	}
}

func generateRow(W int) string {
	s := ""
	if W%7 == 5 {
		s = "55555"
	}
	for w := W - W%7; w > 0; w -= 7 {
		s += "2255555"
	}
	if W%7 == 2 {
		s += "22"
	}
	return s
}

func generateRows(W int, r1, r2 string) (string, string) {
	s1, s2 := "", ""
	if W%7 == 6 {
		s1 = r1
		s2 = r2
	}
	for w := W - W%7; w > 0; w -= 7 {
		s1 += "2" + r1
		s2 += "2" + r2
	}
	if W%7 == 1 {
		s1 += "2"
		s2 += "2"
	}
	return s1, s2
}
