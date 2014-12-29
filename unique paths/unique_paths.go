package main

import (
	"fmt"
	"time"
)

func uniquePaths(m int, n int) int {
	return paths(1, 1, m, n)
}

func paths(row int, col int, m int, n int) int {
	if row > m || col > n {
		return 0
	}
	if row == m && col == n {
		return 1
	}
	return paths(row, col+1, m, n) + paths(row+1, col, m, n)
}

func uniquePathsDp(m int, n int) int {
	arrPath := make([][]int, m)
	lenRow := len(arrPath)
	for i := 0; i < lenRow; i++ {
		arrPath[i] = make([]int, n)
	}
	fmt.Println(arrPath)
	return 0
}

func main() {
	startTime := time.Now().UnixNano()
	num := uniquePathsDp(10, 10)
	endTime := time.Now().UnixNano()
	fmt.Println(num)
	cost := (endTime - startTime) / 1000
	fmt.Printf("cost:%d(ms)", cost)
}
