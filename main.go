package main

import (
	"elp_projet/utils"
	"fmt"
)

func main() {
	a := []string{"apple", " banana", "grap", " cherry"}
	dict := utils.Extraire()
	fmt.Println(utils.Corrections(a, dict))
	utils.New_dic()

}
