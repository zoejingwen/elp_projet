# ELP Project 
## Levenshtein distance to find doublons and match names from various sources
## Création d'un modèle client-serveur TCP
### liste TAF :
* serveur TCP
* texte a corriger à entrer par le client DONE V
* enlever les go routines de la phrase (car met les mots dans le désordre) DONE V
* ajouter go routine dans le dico et adapter la quantité de go routines (paramètre modifiable en fct des capacités du serveur)
* changer le type de sortie de la fonction Corrections (car map met les clés dans un ordre aléatoire) (par exemple remplacer par un liste de tuple : string et liste de string)