async function choisirJoueur(joueurs, dernierJoueur) {
    if (!dernierJoueur) {
        const randomIndex = Math.floor(Math.random() * joueurs.length);
        return joueurs[randomIndex];
    } else {
        let index = joueurs.indexOf(dernierJoueur);
        return joueurs[(index + 1) % joueurs.length];
    }
}
module.exports = choisirJoueur;