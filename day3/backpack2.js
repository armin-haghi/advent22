// 1. parse three lines line 
// 2. find duplicates in entries
// 3. convert to priority a-z 1 to 26 A-Z to 27-52
let totalPrio = 0;
const inputFile = 'input.txt';
const groupSize = 3;

const fs = require('fs');

const data = fs.readFileSync(inputFile, 'UTF-8');
const lines = data.split(/\r?\n/);

function find_duplicate_letter(str1, str2) {
    let reply = "";
    for (var i = 0; i < str2.length; i++) {
        let pos = 0;
        pos = str1.indexOf(str2[i]);
        if (pos > -1) {
            reply = reply + str1[pos];
        } 
    }
    return reply;
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

let totalLines = lines.length -1 ;
let linesInGroup = [];

lines.forEach((line, lineNumber) => {
    linesInGroup.push(line);
    if ((linesInGroup.length == groupSize) || (lineNumber == totalLines)) {
        console.log(linesInGroup, lineNumber, totalLines);

        let firstMatch = find_duplicate_letter(linesInGroup[0], linesInGroup[1]);
        let groupMatch = find_duplicate_letter(firstMatch, linesInGroup[2]);
        console.log(groupMatch);
        let prio = convert_letter_to_priority(groupMatch);
        totalPrio = totalPrio + prio;
        linesInGroup = [];
    } 
  })

  console.log("Total =", totalPrio);

