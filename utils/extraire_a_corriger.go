package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func txt_to_liste(txt_a_corriger string) []string {
	txt_a_corriger = strings.ToLower(txt_a_corriger)
	fmt.Print(txt_a_corriger)

	// On remplace la ponctuation par des espaces
	var sb strings.Builder
	for _, r := range txt_a_corriger {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
			sb.WriteRune(' ')
		} else {
			sb.WriteRune(r)
		}
	}
	fmt.Print(txt_a_corriger)
	txt_a_corriger = sb.String()
	fmt.Print(txt_a_corriger)

	// On sépare le texte en mots
	liste := strings.Fields(txt_a_corriger)
	fmt.Print(txt_a_corriger)

	return liste
}

func Recup_txt_a_corriger() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Entrez le texte à corriger : ")
	txt_a_corriger, _ := reader.ReadString('\n')
	fmt.Print(txt_a_corriger)
	return txt_to_liste(txt_a_corriger)
}
