package utils

import (
	"sync"
)

func Correction(mot string, dict []string) []string {
	res := []string{}
	for i := 0; i < len(dict); i++ {
		distance := Dist_lev(mot, dict[i])
		difference := Pourcentage(mot, dict[i], distance)
		if difference <= 25 {
			res = append(res, dict[i])
		}
		if difference == 0 { // si on trouve un mot identique on arrete la recherche
			res = []string{dict[i]}
			break
		}
	}
	return res
}
func Corrections(a_corriger, dict []string) map[string][]string {

	results := make(map[string][]string)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, mot := range a_corriger {
		wg.Add(1)
		go func(mot string) {
			defer wg.Done()
			corrected := Correction(mot, dict)
			mutex.Lock()
			results[mot] = corrected
			mutex.Unlock()
		}(mot)
	}
	wg.Wait()
	return results
}
