package utils

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func Extraire() []string {
	// Ouvrir le fichier JSON
	file, err := os.Open("Base_de_donnees/dict_anglais.json")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture du fichier : %v", err)
	}
	defer file.Close()

	// Scanner le fichier ligne par ligne (chaque ligne est un dictionnaire JSON)
	scanner := bufio.NewScanner(file)
	words := []string{}

	for scanner.Scan() {
		// Lire chaque ligne
		line := scanner.Text()

		// Lire chaque dictionnaire JSON
		var dict map[string]interface{}
		err := json.Unmarshal([]byte(line), &dict)
		if err != nil {
			log.Printf("Erreur lors du décodage de la ligne : %v", err)
			continue
		}

		// Extraire la valeur associée à "word"
		if word, ok := dict["word"].(string); ok {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
	}
	return words
}

func EcrireListeDansFichierJSON(liste []string, cheminFichier string) {
	// Créer ou ouvrir le fichier
	file, err := os.Create(cheminFichier)
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier : %v", err)
	}
	defer file.Close()

	// Encoder la liste en JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(liste)
	if err != nil {
		log.Fatalf("Erreur lors de l'encodage en JSON : %v", err)
	}
}

func New_dic() {
	// Extraire les mots du fichier JSON
	words := Extraire()

	// Ecrire les mots dans un fichier JSON
	EcrireListeDansFichierJSON(words, "Base_de_donnees/dict_anglais.json")
}
