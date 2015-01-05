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

	return pathsNum(m, n, arrPath)
}

func pathsNum(m, n int, arrPath [][]int) int {
	if m < 1 || n < 1 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	if arrPath[m-1][n-1] != 0 {
		return arrPath[m-1][n-1]
	}

	arrPath[m-2][n-1] = pathsNum(m-1, n, arrPath)
	arrPath[m-1][n-2] = pathsNum(m, n-1, arrPath)
	return arrPath[m-2][n-1] + arrPath[m-1][n-2]
}

func main() {
	fmt.Println("-------------DFS solution------------")
	startTime := time.Now().UnixNano()
	num := uniquePaths(17, 17)
	endTime := time.Now().UnixNano()
	fmt.Printf("paths:%d\n", num)
	cost := (endTime - startTime) / 1000000
	fmt.Printf("cost:%d(ms)\n", cost)

	fmt.Println("-------------DP solution------------")
	startTime = time.Now().UnixNano()
	num = uniquePathsDp(17, 17)
	endTime = time.Now().UnixNano()
	fmt.Printf("paths:%d\n", num)
	cost = (endTime - startTime) / 1000
	fmt.Printf("cost:%d(ms)", cost)

}
