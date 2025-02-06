const { generateSlug } = require("random-word-slugs");
const readline = require("readline");
function word() {
    let word = generateSlug(1);
    return word;
}
module.exports = word;

