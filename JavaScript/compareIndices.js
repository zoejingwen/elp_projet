const readline = require("readline");

async function compareIndices(listeIndices) {
    return new Promise((resolve) => {
        const listeIndicesfiltre = [];
        for (let i = 0; i < listeIndices.length; i++) {
            if (!(listeIndicesfiltre.includes(listeIndices[i]))) {
                listeIndicesfiltre.push(listeIndices[i]);
            }
            else {
                for (let j = 0; j < listeIndicesfiltre.length; j++){
                    if (listeIndices[i] == listeIndicesfiltre[j]) {
                        listeIndicesfiltre.splice(j,1)
                    }
                }
            }
        }
        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout,
        });
        rl.question(`quels indices ne sont pas valides?`, (StringInvalide) => {
            let listInvalide = StringInvalide.split(" ");
            for (let i = 0; i < listInvalide.length; i++) {
                for (let j = 0; j < listeIndicesfiltre.length; j++){
                    if (listInvalide[i] == listeIndicesfiltre[j]){
                        listeIndicesfiltre.splice(j, 1);
                        j--; //ajuste l'index après suppression
                    }
                }       
            }
            console.log(listeIndicesfiltre);
            rl.close();

            resolve(listeIndicesfiltre);
        });
    });
}

module.exports = compareIndices;
