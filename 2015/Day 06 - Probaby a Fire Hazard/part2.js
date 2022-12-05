const { readFileSync } = require('fs')

const lights = Array.from({ length: 1000 }, () =>
  Array.from({ length: 1000 }, () => 0)
)

function setLights(action, x1, y1, x2, y2) {
  for (let x = x1; x <= x2; x++) {
    for (let y = y1; y <= y2; y++) {
      lights[x][y] +=
        action === 'turn on' ? 1 : action === 'toggle' ? 2 : lights[x][y] > 0 ? -1 : 0
    }
  }
}

readFileSync('input.txt', 'utf8')
  .split('\n')
  .map((instruction) => {
    const [_, action, c1x, c1y, c2x, c2y] = instruction.match(
      /(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)/
    )
    setLights(action, +c1x, +c1y, +c2x, +c2y)
  })

console.log(
  lights.reduce((lit, row) => lit + row.reduce((lit, light) => lit + light, 0), 0)
)