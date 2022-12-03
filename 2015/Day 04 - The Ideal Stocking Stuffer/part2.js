const { readFileSync } = require('fs')
const crypto = require('crypto')

const secretKey = readFileSync('input.txt', 'utf8').trim();
const startsWithSixLeadingZeroes = (hash) => hash.slice(0, 6) === '000000';

let i = 0;
while (true) {
  const hash = crypto.createHash('md5').update(secretKey + i).digest('hex');
  if (startsWithSixLeadingZeroes(hash)) {
    //console.log(`${i} is the lowest positive number that produces a hash that starts with 6 leading zeroes (${hash})`);
    console.log(i);
    break;
  }
  i++;
}
