package main

import (
	"ELP_PROJET/utils"
	"fmt"
)

func main() {
	a := "apple banana grape cherry"
	dict := utils.Extraire()
	fmt.Println(utils.Corrections(a, dict))

}
