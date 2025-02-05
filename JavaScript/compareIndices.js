import readline from 'node:readline'

function compareIndices(listeIndices) {
    return new Promise((resolve) => {
        const listeIndicesfiltre = [];
        //console.log(!(listeIndicesfiltre.includes(listeIndices[2])))
        for (let i = 0; i < listeIndices.length; i++) {
            if (!(listeIndicesfiltre.includes(listeIndices[i]))) {
                listeIndicesfiltre.push(listeIndices[i]);
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

(async () => {
    const result = await compareIndices(["a", "a", "b"]);
    console.log("Résultat final :", result);
})();