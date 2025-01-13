package main

import "fmt"

var a string = "kitten"

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
func dist_lev(a, b string) int {
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

func pourcentage(a string, b string, dist int) float64 {
	var res float64
	a_rune := []rune(a)
	b_rune := []rune(b)
	n, m := len(a_rune), len(b_rune)
	res = float64(dist) / float64(max(n, m)) * 100
	return res
}

func main() {
	a := "ogiyowiiiigg <,?[!12]"
	b := "ogiyowiiiggs <,?[!12]"
	fmt.Println(dist_lev(a, b))
	fmt.Printf("la différence entre les deux chaînes est %v %%\n", pourcentage(a, b, dist_lev(a, b)))
}
