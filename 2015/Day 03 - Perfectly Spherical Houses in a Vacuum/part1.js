const { readFileSync } = require('fs')

let lastX = (lastY = 0)

console.log(
  readFileSync('input.txt', 'utf8')
    .split('')
    .reduce(
      (houses, direction) => {
        const x = lastX + (direction === '>' ? 1 : direction === '<' ? -1 : 0)
        const y = lastY + (direction === '^' ? 1 : direction === 'v' ? -1 : 0)
        if (
          houses.filter((house) => house.x === x && house.y === y).length === 0
        ) {
          houses.push({ x, y })
        }
        lastX = x
        lastY = y
        return houses
      },
      [{ x: 0, y: 0 }]
    )
    .length
)
