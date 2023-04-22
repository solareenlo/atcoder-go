package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	if n == 1 {
		var x int
		fmt.Fscan(in, &x)
		var q int
		fmt.Fscan(in, &q)
		for i := 0; i < q+1; i++ {
			fmt.Fprintln(out, 0)
		}
		return
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	var dp [3030][3030]int
	for i := 0; i < n; i++ {
		j := i - 1
		used := make([]int, n)
		for j >= 0 && a[j] != a[i] {
			if used[a[j]] == 0 {
				if a[i] < a[j] {
					dp[a[i]][a[j]] += 1
				}
				if a[i] > a[j] {
					dp[a[j]][a[i]] += 1
				}
			}
			used[a[j]] = 1
			j--
		}
		for j := 0; j < n; j++ {
			if a[i] < j {
				dp[a[i]][j] = max(dp[a[i]][j], 1)
			}
			if a[i] > j {
				dp[j][a[i]] = max(dp[j][a[i]], 1)
			}
		}
	}
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			cnt[dp[j][i]] += 1
		}
	}
	ans := 0
	for i := 0; i < n+1; i++ {
		if cnt[i] != 0 {
			ans = i
		}
	}
	fmt.Fprintln(out, n-ans)

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		l := make([]int, n)
		r := make([]int, n)
		li := x - 1
		ri := x + 1
		for li >= 0 && a[li] != a[x] {
			l[a[li]] = 1
			li--
		}
		for ri < n && a[ri] != a[x] {
			r[a[ri]] = 1
			ri++
		}
		if li == -1 {
			if ri == n {
				for j := 0; j < n; j++ {
					if a[x] < j {
						cnt[dp[a[x]][j]]--
						dp[a[x]][j] = 0
					}
					if a[x] > j {
						cnt[dp[j][a[x]]]--
						dp[j][a[x]] = 0
					}
				}
			} else {
				for j := 0; j < n; j++ {
					if l[j] != 0 && r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] -= 2
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] -= 2
							cnt[dp[j][a[x]]]++
						}
					} else if r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] -= 1
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] -= 1
							cnt[dp[j][a[x]]]++
						}
					}
				}
			}
		} else {
			if ri == n {
				for j := 0; j < n; j++ {
					if l[j] != 0 && r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] -= 2
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] -= 2
							cnt[dp[j][a[x]]]++
						}
					} else if l[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] -= 1
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] -= 1
							cnt[dp[j][a[x]]]++
						}
					}
				}
			} else {
				for j := 0; j < n; j++ {
					if l[j] != 0 && r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] -= 2
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] -= 2
							cnt[dp[j][a[x]]]++
						}
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			l[i] = 0
			r[i] = 0
		}
		li = x - 1
		ri = x + 1
		a[x] = y
		for li >= 0 && a[li] != a[x] {
			l[a[li]] = 1
			li--
		}
		for ri < n && a[ri] != a[x] {
			r[a[ri]] = 1
			ri++
		}
		if li == -1 {
			if ri == n {
				for j := 0; j < n; j++ {
					if a[x] < j {
						dp[a[x]][j] = 1 + l[j] + r[j]
						cnt[dp[a[x]][j]]++
					}
					if a[x] > j {
						dp[j][a[x]] = 1 + l[j] + r[j]
						cnt[dp[j][a[x]]]++
					}
				}
			} else {
				for j := 0; j < n; j++ {
					if l[j] != 0 && r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] += 2
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] += 2
							cnt[dp[j][a[x]]]++
						}
					} else if r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] += 1
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] += 1
							cnt[dp[j][a[x]]]++
						}
					}
				}
			}
		} else {
			if ri == n {
				for j := 0; j < n; j++ {
					if l[j] != 0 && r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] += 2
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] += 2
							cnt[dp[j][a[x]]]++
						}
					} else if l[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] += 1
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] += 1
							cnt[dp[j][a[x]]]++
						}
					}
				}
			} else {
				for j := 0; j < n; j++ {
					if l[j] != 0 && r[j] != 0 {
						if a[x] < j {
							cnt[dp[a[x]][j]]--
							dp[a[x]][j] += 2
							cnt[dp[a[x]][j]]++
						}
						if a[x] > j {
							cnt[dp[j][a[x]]]--
							dp[j][a[x]] += 2
							cnt[dp[j][a[x]]]++
						}
					}
				}
			}
		}
		ans := 0
		for i := 0; i < n+1; i++ {
			if cnt[i] > 0 {
				ans = i
			}
		}
		fmt.Fprintln(out, n-ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
