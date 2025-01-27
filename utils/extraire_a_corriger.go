package utils

import (
	"strings"
	"unicode"
)

func Txt_to_liste(txt_a_corriger string) []string {
	txt_a_corriger = strings.ToLower(txt_a_corriger)
	// On remplace la ponctuation par des espaces
	var sb strings.Builder
	for _, r := range txt_a_corriger {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
			sb.WriteRune(' ')
		} else {
			sb.WriteRune(r)
		}
	}
	txt_a_corriger = sb.String()

	// On sépare le texte en mots
	liste := strings.Fields(txt_a_corriger)

	return liste
}
