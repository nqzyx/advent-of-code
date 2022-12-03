const { readFileSync } = require('fs');

let position = 0;

readFileSync('./input.txt', 'utf-8')
.split('')
.reduce((floor, direction, index) => {
  if (direction === '(') {
    floor++;
  } else if (direction === ')') {
    floor--;
  }
  if (floor === -1 && position === 0) {
    position = index + 1;
  }
  return floor;
}, 0);

console.log(position);
