package utils

import (
	"runtime"
	"sync"
)

func Correction(mot string, dict []string) []string {
	res := []string{}

	var mutex sync.Mutex
	var wait_group sync.WaitGroup
	poolSize := runtime.NumCPU()
	pool := make(chan struct{}, poolSize)
	stop := make(chan struct{})
	for i := 0; i < len(dict); i++ {
		mot_dico := dict[i]
		wait_group.Add(1)
		pool <- struct{}{} //bloque si autant de go routine que de coeurs
		go func(mot_dico string) {
			defer wait_group.Done()
			select {
			case <-stop:
				<-pool
				return //arrete la go routine si un mot identique a été trouvé
			default:
				distance := Dist_lev(mot, mot_dico)
				difference := Pourcentage(mot, mot_dico, distance)
				if difference <= 25 {
					mutex.Lock()
					res = append(res, mot_dico)
					mutex.Unlock()
				}
				if difference == 0 {
					close(stop)
					mutex.Lock()
					res = []string{mot_dico}
					mutex.Unlock()
				}
			}
			<-pool
		}(mot_dico)
	}
	wait_group.Wait()
	return res
}

func Corrections(a_corriger, dict []string) map[string][]string {

	results := make(map[string][]string)
	for _, mot := range a_corriger {
		corrected := Correction(mot, dict)
		results[mot] = corrected
	}
	return results
}
