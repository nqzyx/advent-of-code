// AoC 2022 Day 14 - Regolith Reservoir - Part 1

import { readFileSync } from 'fs'

import { create, all, max } from 'mathjs'
const config = {}
const math = create(all, config)

const sandSource = [500, 0]

const emptySymbol = '.'
const rockSymbol = '#'
const sandFallingSymbol = 'v'
const sandSymbol = 'o'
const sandSourceSymbol = '+'

let [minY, minX] = sandSource
let [maxY, maxX] = sandSource

const getX = (x) => x - minX
const getY = (y) => y - minY

const getLoc = (x, y) => objectMap[getX(x)][getY(y)]

const setLoc = (x, y, value) => {
  objectMap[getX(x)][getY(y)] = value
}

const instructions = readFileSync('input.txt', 'utf8')
  .split('\n')
  .map((i) => i.split(' -> ').map((pt) => eval(`[${pt}]`)))
//.map((pt) => math.subtract(pt, 1))

// console.log(instructions)
// process.exit(0)

instructions.forEach((i) => {
  i.forEach((v) => {
    const [vY, vX] = v
    minY = math.min(minY, vY)
    maxY = math.max(maxY, vY)
    minX = math.min(minX, vX)
    maxX = math.max(maxX, vX)
  })
})

// console.log({ yMin: minY, yMax: maxY, xMin: minX, xMax: maxX })
// process.exit(0)

const objectMap = math
  .range(getX(minX), getX(maxX), true)
  .map(() => math.range(getY(minY), getY(maxY), true).map(() => emptySymbol))
  .toArray()

setLoc(...(sandSource.slice().reverse()), sandSourceSymbol)

// objectMap.forEach((r, x) => console.log(minX + x, ':', r.join('')))
// process.exit(0)

instructions.forEach((i) => {
  let lastV
  i.forEach((v, vi) => {
    if (vi > 0) {
      math.range(math.min(v[1], lastV[1]), math.max(v[1], lastV[1]), true)
        .forEach((x) =>
          math.range(math.min(v[0], lastV[0]), math.max(v[0], lastV[0]), true)
            .forEach((y) => setLoc(x, y, rockSymbol))
        )
    }
    lastV = v
  })
})

// objectMap.forEach((r, x) => console.log(minX + x, ':', r.join('')))
// process.exit(0)

let done = false
let sandUnits = 0

while (!done) {
  let loc = sandSource
  let falling = true

  while (falling) {
    const [y, x] = loc

    const oldSymbol = getLoc(x, y)
    setLoc(x, y, sandFallingSymbol)

    // Off the bottom of the board
    if (x + 1 > maxX) {
        falling = false
        done = true      
        break
    }

    // Fall straight down
    if (getLoc(x + 1, y) === emptySymbol) {
      setLoc(x, y, oldSymbol)
      loc = [y, x + 1]
      continue
    }

    // Off the left of the board
    if (y - 1 < minY) {
        falling = false
        done = true      
        break
    }

    // Fall down to left
    if (getLoc(x + 1, y - 1) === emptySymbol) {
      setLoc(x, y, oldSymbol)
      loc = [y - 1, x + 1]
      continue
    }

    // Off the right of the board
    if (y + 1 > maxY) {
        falling = false
        done = true      
        break
    }

    // Fall down to right
    if (getLoc(x + 1, y + 1) === emptySymbol) {
      setLoc(x, y, oldSymbol)
      loc = [y + 1, x + 1]
      continue
    }

    // Stop falling
    setLoc(x, y, sandSymbol)
    sandUnits += 1
    falling = false
}

// console.log('===', 'State after', sandUnits, 'Sand Units', sandUnits, '===')
// objectMap.forEach((r, x) => console.log(minX + x, ':', r.join('')))

  if (falling) break
}

// console.log('===', 'State after', sandUnits, 'Sand Units', sandUnits, '===')
// objectMap.forEach((r, x) => console.log(minX + x, ':', r.join('')))

console.log('Answer:', sandUnits)
// count the sand...
