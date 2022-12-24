// 1. parse each line 
// 2. find length / 2 per line = halflenght
// 3. split each line at halflenght
// 4. loop for each letter in str1 and str2 until duplicate is found


let str = 'vJrwpWtwJgWrhcsFMMfFFhFp';

function find_duplicate_letter (str: string) {
  let split_lenght = str.length/2;
  let str1 = str.substring(0,split_lenght);
  let str2 = str.substring(split_lenght,split_lenght*2);
  for (var i = 0; i < split_lenght; i++) {
      let pos: number;
      pos = str2.indexOf(str1[i]);
      if (pos > -1) {
          return str2[pos];
      }
    }
}

console.log(find_duplicate_letter(str))