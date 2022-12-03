const { readFileSync } = require('fs')

let last = { x: 0, y: 0 }

function visitedHouses(dirs, first = { x: 0, y: 0 }) {
  last = first
  return dirs.reduce(
    (vh, d) => {
      const next = {
        x: last.x + (d === '>' ? 1 : d === '<' ? -1 : 0),
        y: last.y + (d === '^' ? 1 : d === 'v' ? -1 : 0),
      }
      if (vh.filter((h) => h.x === next.x && h.y === next.y).length === 0) {
        vh.push(next)
      }
      last = next
      return vh
    },
    [first]
  )
}

const dirs = readFileSync('input.txt', 'utf8').split('')

const housesVisitedBySanta = visitedHouses(dirs.filter((_, i) => i % 2 === 0));
const housesVisitedByRoboSanta = visitedHouses(dirs.filter((_, i) => i % 2 === 1));

const totalVisitedHouses = housesVisitedBySanta.concat(
  housesVisitedByRoboSanta.filter(
    (houseVisitedByRoboSanta) =>
      housesVisitedBySanta.find(
        (houseVisitedBySanta) =>
          houseVisitedBySanta.x === houseVisitedByRoboSanta.x &&
          houseVisitedBySanta.y === houseVisitedByRoboSanta.y
      ) === undefined
  )
)

console.log(totalVisitedHouses.length)
