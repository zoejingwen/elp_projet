const readline = require('readline-sync');

function reponse(secret_word) {
    let reponse = readline.question("Tu penses que cela doit être quel mot ? ");
    let comparaison;

    if (reponse !== null && reponse !== "") {
        if (reponse === secret_word) {
            comparaison = "réussite";
        } else {
            comparaison = "échec";
        }
    } else {
        comparaison = "pass";
    }

    return comparaison;
}

module.exports = reponse;
