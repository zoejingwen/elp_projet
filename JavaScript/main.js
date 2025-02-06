const choisirJoueur = require("./joueur_actif");
const word = require("./word");
const score = require("./calcul_score");
const compareIndices = require("./compareIndices");
const reponse = require("./reponse");
const getIndices=require("./get_indice");
async function main() {
    let joueurs = ["Alice", "Bob", "Charlie", "David"];
    let total_score = 0;
    let dernierJoueur = null;
    let card_box = [];
    
    for (let i = 0; i < 7; i++) {
        card_box.push(await word());
    }
    
    while (card_box.length > 0) {
        console.log("\nNouvelle manche !");
        let joueur = await choisirJoueur(joueurs, dernierJoueur);
        dernierJoueur = joueur;
        console.log(`Le joueur qui doit deviner est : ${joueur}`);

        let mot_secret = card_box.shift();
        
        const bon = await new Promise((resolve) => {
            rl.question(`Est-ce que vous pouvez bien comprendre ce mot mystère ? (${mot_secret}) (Oui/Non) `, (reponse) => {
                resolve(reponse.toLowerCase());
            });
        });
        
        if (bon === "non") {
            mot_secret = card_box.shift();
            console.log(`Nouveau mot secret : ${mot_secret}`);
        }

        let indices = await getIndices(joueurs.filter(j => j !== joueur));
        let indices_final = await compareIndices(indices);
        console.log(`Indices valides : ${indices_final.join(", ")}`);

        let resultat = await reponse(mot_secret);
        let resultatScore = score(total_score, resultat, card_box);
        total_score = resultatScore.score;
        card_box = resultatScore.card;
        console.log(`Résultat : ${resultat}, Score total : ${total_score}`);
        
        let continuer = await new Promise((resolve) => {
            rl.question("Voulez-vous continuer ? (Oui/Non) ", (reponse) => {
                resolve(reponse.toLowerCase() === "oui");
            });
        });

        if (!continuer) break;
    }
    console.log("Fin du jeu ! Score final :", total_score);
    rl.close();
}

main();