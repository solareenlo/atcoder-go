package main

import "fmt"

func main() {
	var y, m int
	fmt.Scan(&y, &m)
	tar := (y-2014)*12 + m
	month := 13
	year := 2013
	for tar >= month {
		year++
		tar -= month
		month++
	}
	fmt.Println(year, 12+tar)
}
