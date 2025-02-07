function score(score, reponse, card) {
    let newCard = [...card];  // Crée une copie pour éviter la mutation

    if (reponse === "réussite") {
        score += 1;
    } else if (reponse === "pass") {
        // On ne fait rien ici, on ignore la condition
    } else {
        if (newCard.length > 0) {
            newCard.pop();  // Retire un élément en évitant d'éventuels problèmes avec un tableau vide
        }
    }

    return { score: score, card: newCard };  // Retourne un nouvel objet
}

module.exports = score;
