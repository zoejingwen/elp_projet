const { generateSlug } = require("random-word-slugs");
const readline = require("readline");

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function word() {
    let tmp_word = generateSlug(1);
    console.log(`Mot généré : ${tmp_word}`);
    rl.question("Est-ce que vous pouvez bien comprendre ce mot mystère ? (${tmp_word}) (Oui/Non) ", (bon) => {
        if (bon.toLowerCase() === "non") {
            tmp_word = generateSlug(1);
            console.log(`Nouveau mot : ${tmp_word}`);
        } else {
            console.log(`Mot choisi : ${tmp_word}`);
        }
        rl.close();
    });
}

word();


