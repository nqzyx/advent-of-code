const { readFileSync } = require('fs')

function intersection(setA, setB) {
  const _intersection = new Set()
  for (const elem of setB) {
    if (setA.has(elem)) {
      _intersection.add(elem)
    }
  }
  return _intersection
}

console.log(readFileSync('./input.txt', 'utf-8')
  .split('\n')
  .map((sackContents) => {
    const c1 = new Set(sackContents.slice(0, sackContents.length / 2).split(''))
    const c2 = new Set(sackContents.slice(sackContents.length / 2).split(''))
    const d = intersection(c1, c2)
    if (d.size !== 1) {
      throw new Error(`${d.size}: duplicate items: ${d}`)
    }
    const i = [...d][0]
    const p =
      i.toUpperCase() === i
        ? i.charCodeAt(0) - 'A'.charCodeAt(0) + 27
        : i.charCodeAt(0) - 'a'.charCodeAt(0) + 1
    return p
  })
  .reduce((score, p) => score + p, 0)
);
