const { readFileSync } = require('fs');

console.log(readFileSync('./input.txt', 'utf-8')
  .split('')
  .reduce((floor, direction) => {
    if (direction === '(') {
      floor++;
    } else if (direction === ')') {
      floor--;
    }
    return floor;
  }, 0));
