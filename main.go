package main

import (
	"elp_projet/utils"
	"fmt"
)

func main() {
	dict := utils.ExtraireListe("Base_de_donnees/dict_anglais.json")
	a_corriger := utils.ExtraireListe("Base_de_donnees/a_corriger.json")
	fmt.Println(utils.Corrections(a_corriger, dict))

}
