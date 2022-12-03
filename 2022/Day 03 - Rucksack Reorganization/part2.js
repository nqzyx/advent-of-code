const { readFileSync } = require('fs')

function intersection(setA, setB, setC) {
  const _intersection = new Set()
  for (const elem of setA) {
    if (setB.has(elem) && setC.has(elem)) {
      _intersection.add(elem)
    }
  }
  return _intersection
}

console.log(
  readFileSync('./input.txt', 'utf-8')
    .split('\n')
    .reduce((score, _, e, elves) => {
      if (e % 3 !== 0) return score
      const d = intersection(
        new Set(elves[e].split('')),
        new Set(elves[e + 1].split('')),
        new Set(elves[e + 2].split(''))
      )
      if (d.size !== 1) {
        throw new Error(`${d.size} duplicate items: ${d}`)
      }
      const i = [...d][0]
      return (
        score +
        i.charCodeAt(0) -
        (i === i.toUpperCase() ? 'A'.charCodeAt(0) - 27 : 'a'.charCodeAt(0) - 1)
      )
    }, 0)
)
