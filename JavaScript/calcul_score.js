function score(total_score, reponse, card) {
    if (reponse === "reussite") {
        total_score = score+1;
    }
    else if (reponse === "pass") {
        // On ne fait rien ici, on ignore la condition
    }
    else {
        card.pop();  // On enlève le dernier élément de la carte
    }

    return { score: total_score, card: card };  // Retourne un objet avec les deux valeurs
}
module.exports = score;
