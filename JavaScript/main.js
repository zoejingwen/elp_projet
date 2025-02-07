const choisirJoueur = require("./joueur_actif");
const word = require("./select_secret_word");
const score = require("./calcul_score");
const compareIndices = require("./compareIndices");
const reponse = require("./reponse");
const readline = require("readline");
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
async function getIndices(joueurs) {
    return new Promise((resolve) => {
        let indices = [];
        let count = 0;
        
        function askNext() {
            if (count >= joueurs.length) {
                resolve(indices);
                return;
            }
            rl.question(`${joueurs[count]} donne un indice : `, (indice) => {
                indices.push(indice);
                count++;
                askNext();
            });
        }
        askNext();
    });
}
async function poserQuestion(question) {
    return new Promise((resolve) => {
        rl.question(question, (reponse) => {
            resolve(reponse.trim().toLowerCase());
        });
    });
}
async function main() {
    let joueurs = ["Alice", "Bob", "Charlie","Tom","Lisa"];
    let total_score = 0;
    let dernierJoueur = null;
    let card_box = [];
    
    for (let i = 0; i < 7; i++) {
        let mot = word();
        card_box.push(mot);
    }
    
    
    while (card_box.length > 0) {
        console.log("\nNouvelle manche !");
        let joueur = await choisirJoueur(joueurs, dernierJoueur);
        dernierJoueur = joueur;
        console.log(`Le joueur qui doit deviner est : ${joueur}`);

        let mot_secret = card_box.shift();
        console.log(mot_secret);
        let bon = await poserQuestion(`Est-ce que vous pouvez bien comprendre ce mot mystère ? (${mot_secret}) (Oui/Non) `);
        
        if (bon === "non") {
            mot_secret = card_box.shift();
            console.log(`Nouveau mot secret : ${mot_secret}`);
        }

        let indices = await getIndices(joueurs.filter(j => j !== joueur));
        let indices_final = await compareIndices(indices);
        console.log(`Indices valides : ${indices_final.join(", ")}`);

        let resultat =reponse(mot_secret);
        let resultatScore = score(total_score, resultat, card_box);
        total_score = resultatScore.score;
        card_box = resultatScore.card;
        console.log(`Résultat : ${resultat}, Score total : ${total_score},amount of card:${card_box.length}`);
        let continuer = await poserQuestion("Voulez-vous continuer ? (Oui/Non) ");
        if (continuer === "non") {
            break;
        }

    }
    console.log("Attente de la réponse...")
    console.log("Fin du jeu ! Score final :", total_score);
    rl.close();
}

main();
