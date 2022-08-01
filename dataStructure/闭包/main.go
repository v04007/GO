package main

import "fmt"

func Rescuvie(n int) (result int) {
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}

//{5 * Rescuvie(4)}
//{5 * {4 * Rescuvie(3)}}
//{5 * {4 * {3 * Rescuvie(2)}}}
//{5 * {4 * {3 * {2 * Rescuvie(1)}}}}
//{5 * {4 * {3 * {2 * 1}}}}
//{5 * {4 * {3 * 2}}}
//{5 * {4 * 6}}
//{5 * 24}
//120

func sum(n int) int {
	total := 0
	// 从1加到N, 1+2+3+4+5+..+N
	for i := 1; i <= n; i++ {
		total = total + i
	}
	return total
}
