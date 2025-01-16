package main

import (
	"elp_projet/utils"
	"fmt"
)

func main() {
	a_corriger := utils.Recup_txt_a_corriger()
	fmt.Println(a_corriger)
	dict := utils.ExtraireListe("Base_de_donnees/dict_anglais.json")
	fmt.Println(utils.Corrections(a_corriger, dict))

}
