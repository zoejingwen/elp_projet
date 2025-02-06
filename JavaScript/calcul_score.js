function score(score, reponse, card) {
    if (reponse === "reussite") {
        score += 1;
    }
    else if (reponse === "pass") {
        // On ne fait rien ici, on ignore la condition
    }
    else {
        card.pop();  // On enlève le dernier élément de la carte
    }

    return { score: score, card: card };  // Retourne un objet avec les deux valeurs
}
