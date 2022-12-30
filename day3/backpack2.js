// 1. parse each line 
// 2. find length / 2 per line = halflenght
// 3. split each line at halflenght
// 4. loop for each letter in str1 and str2 until duplicate is found
// 5. convert to priority a-z 1 to 26 A-Z to 27-52
let totalPrio = 0;
let inputFile = 'input.txt';

const fs = require('fs');

const data = fs.readFileSync(inputFile, 'UTF-8');
const lines = data.split(/\r?\n/);

function find_duplicate_letter(str) {
    let split_lenght = str.length / 2;
    let str1 = str.substring(0, split_lenght);
    let str2 = str.substring(split_lenght, split_lenght * 2);
    for (var i = 0; i < split_lenght; i++) {
        let pos = void 0;
        pos = str2.indexOf(str1[i]);
        if (pos > -1) {
            return str2[pos];
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

