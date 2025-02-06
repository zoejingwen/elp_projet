const { generateSlug } = require("random-word-slugs");
const readline = require("readline");

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function word() {
    let tmp_word = generateSlug(1);
    console.log(`Mot secret: ${tmp_word}`);
}
module.exports = word;

