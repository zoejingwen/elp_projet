package utils

func Pourcentage(a string, b string, dist int) float64 {
	var res float64
	a_rune := []rune(a)
	b_rune := []rune(b)
	n, m := len(a_rune), len(b_rune)
	res = float64(dist) / float64(max(n, m)) * 100
	return res
}
