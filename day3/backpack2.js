// 1. parse three lines line 
// 2. find duplicates in entries
// 3. convert to priority a-z 1 to 26 A-Z to 27-52
let totalPrio = 0;
let inputFile = 'input.txt';

const fs = require('fs');

const data = fs.readFileSync(inputFile, 'UTF-8');
const lines = data.split(/\r?\n/);

function find_duplicate_letter(str1, str2) {
    for (var i = 0; i < str2.lenght; i++) {
        let pos = void 0;
        pos = str1.indexOf(str2[i]);
        if (pos > -1) {
            return str1[pos];
        }
    }
}

function convert_letter_to_priority(str) {
    let prio;
    let base = str.charCodeAt(0);
    if (base < 97) {
        prio = base - 64 + 26;
    } else {
        prio = base - 96;
    }
    return prio;
}

lines.forEach(line => {
    let str = find_duplicate_letter(line);
    let prio = convert_letter_to_priority(str);
    console.log(str, prio);
    totalPrio = totalPrio + prio;
  })

  console.log("Total =", totalPrio);

