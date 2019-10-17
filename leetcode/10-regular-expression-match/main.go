package main

import "fmt"

func main() {
	dp := make([][]bool, 5)
	for i := range dp {
		dp[i] = make([]bool, 3)
	}
	for i := range dp {
		for j := range dp[i] {
			fmt.Println(i, j, dp[i][j])
		}
	}
}
