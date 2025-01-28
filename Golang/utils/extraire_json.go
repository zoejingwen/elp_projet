package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ExtraireListe(cheminFichier string) []string {
	// Ouvrir le fichier JSON
	file, err := os.Open(cheminFichier)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture du fichier : %v", err)
	}
	defer file.Close()

	// Lire le contenu du fichier
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
	}

	// Décoder le JSON
	var words []string
	err = json.Unmarshal(byteValue, &words)
	if err != nil {
		log.Fatalf("Erreur lors du décodage du JSON : %v", err)
	}

	return words
}
