package main

import (
	"fmt"
	"strconv"
)

func main() {
	var hour, minute, second int
	fmt.Scan(&hour, &minute, &second)
	var diffSecond int
	fmt.Scan(&diffSecond)

	var clock Clock
	clock.set(hour, minute, second)
	fmt.Println(clock.toStr())
	clock.shift(diffSecond)
	fmt.Println(clock.toStr())
}

type Clock struct {
	hour, minute, second int
}

func (c *Clock) set(h, m, s int) {
	c.hour = h
	c.minute = m
	c.second = s
}

func (c Clock) toStr() string {
	h := int2str(c.hour)
	m := int2str(c.minute)
	s := int2str(c.second)
	return h + ":" + m + ":" + s
}

func (c *Clock) shift(diffSecond int) {
	diff_h := diffSecond / 3600
	diff_m := (diffSecond % 3600) / 60
	diff_s := (diffSecond % 3600) % 60
	setMinusTime(&c.second, diff_s, &c.minute)
	setMinusTime(&c.minute, diff_m, &c.hour)
	if c.hour+diff_h < 0 {
		c.hour += 24
		c.hour += diff_h
	} else {
		c.hour += diff_h
	}
	roundUp(&c.second, &c.minute)
	roundUp(&c.minute, &c.hour)
	if c.hour == 24 {
		c.hour = 0
	}
}

func roundUp(time1, time2 *int) {
	if *time1 >= 60 {
		*time1 = *time1 % 60
		*time2 += 1
	}
}

func setMinusTime(time1 *int, diff int, time2 *int) {
	if *time1+diff < 0 {
		*time2 -= 1
		*time1 += 60
		*time1 += diff
	} else {
		*time1 += diff
	}
}

func int2str(num int) string {
	var tmp string
	if num < 10 && num >= 0 {
		tmp = "0" + strconv.Itoa(num)
	} else {
		tmp = strconv.Itoa(num)
	}
	return tmp
}
