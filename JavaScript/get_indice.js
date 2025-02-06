async function getIndices(joueurs) {
    return new Promise((resolve) => {
        let indices = [];
        let count = 0;
        
        function askNext() {
            if (count >= joueurs.length - 1) {
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
module.exports = getIndices;