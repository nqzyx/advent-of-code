const { readFileSync } = require('fs')

let last = { x: 0, y: 0 };

console.log(
  readFileSync('input.txt', 'utf8')
    .split('')
    .reduce(
      (houses, direction) => {
        const next = {
          x: last.x + (direction === '>' ? 1 : direction === '<' ? -1 : 0),
          y: last.y + (direction === '^' ? 1 : direction === 'v' ? -1 : 0)
        };
        if (
          houses.filter((house) => house.x === next.x && house.y === next.y).length === 0
        ) {
          houses.push(next)
        }
        last = next;
        return houses;
      },
      [last]
    )
    .length
)
