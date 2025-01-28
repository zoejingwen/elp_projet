package utils

func min(a, b, c int) int {

	switch {
	case a < b && a < c:
		return a
	case b < a && b < c:
		return b
	default:
		return c
	}
}
func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func Dist_lev(a, b string) int {
	a_rune := []rune(a)
	b_rune := []rune(b)
	n, m := len(a_rune), len(b_rune)
	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			switch {
			case i == 0:
				dist[i][j] = j
			case j == 0:
				dist[i][j] = i
			case a[i-1] == b[j-1]:
				dist[i][j] = dist[i-1][j-1]
			default:
				dist[i][j] = 1 + min(dist[i-1][j], dist[i][j-1], dist[i-1][j-1])
			}
		}
	}
	return dist[n][m]
}
