package study

import (
	"demo/pkg"
)

func EmptyToByte() int {
	s := ""
	return len([]byte(s))
}

func LevCompare() bool {
	a := "ABC"
	b := "ABCD"
	distance := pkg.ComputeDistance(a, b)
	return distance == Levenshtein(a, b)
}

func Levenshtein(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	dp := make([][]int, lenA+1)
	for i := range dp {
		dp[i] = make([]int, lenB+1)
	}

	for i := 0; i <= lenA; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		dp[0][j] = j
	}

	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			prev := min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			if a[i-1] != b[j-1] {
				prev += 1
			}
			dp[i][j] = prev
		}
	}

	return dp[lenA-1][lenB-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
