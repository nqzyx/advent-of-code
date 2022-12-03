const { readFileSync } = require('fs')
const crypto = require('crypto')

const secretKey = readFileSync('input.txt', 'utf8').trim();
const startsWithFiveLeadingZeroes = (hash) => hash.slice(0, 5) === '00000';

let i = 0;
while (true) {
  const hash = crypto.createHash('md5').update(secretKey + i).digest('hex');
  if (startsWithFiveLeadingZeroes(hash)) {
    //console.log(`${i} is the lowest positive number that produces a hash that starts with 5 leading zeroes (${hash})`);
    console.log(i);
    break;
  }
  i++;
}
