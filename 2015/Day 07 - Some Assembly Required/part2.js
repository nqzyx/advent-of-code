const { readFileSync } = require('fs')

function wireCheck (w) {
  return `('${w}' in solvedWires)`
}

function wireValue (w) {
  return `solvedWires['${w}']`
}

const ops = { AND: '&', OR: '|', LSHIFT: '<<', RSHIFT: '>>' }

const solvedWires = {}
let unsolvedWires = {}

readFileSync('part2_input.txt', 'utf8')
  .split('\n')
  .forEach((connection) => {
    const [src, w] = connection.split(' -> ')
    if (src.match(/^\d+$/)) { // Input is a constant
      solvedWires[w] = parseInt(src)
    }
    if (src.match(/^[a-z]+$/)) { // Input is another wire
      unsolvedWires[w] = `!${wireCheck(src)} ? undefined : ${wireValue(src)}`
    }
    if (src.match(/^NOT ([a-z]+)$/)) { // Input is a NOT gate
      const [, iw] = src.split(' ')
      unsolvedWires[w] = `!${wireCheck(iw)} ? undefined : (~${wireValue(iw)} & 0xffff)`
    }
    if (src.match(/ (AND|OR|LSHIFT|RSHIFT) /)) { // Input is an AND gate
      const [l, o, r] = src.match(/^(\d+|[a-z]+) ([A-Z]+) (\d+|[a-z]+)$/).slice(1, 4)
      const op = ops[o]
      const lCheck = (l.match(/^[a-z]+$/)) ? wireCheck(l) : 'true'
      const rCheck = (r.match(/^[a-z]+$/)) ? wireCheck(r) : 'true'
      const lValue = (l.match(/^[a-z]+$/)) ? wireValue(l) : l
      const rValue = (r.match(/^[a-z]+$/)) ? wireValue(r) : r
      unsolvedWires[w] = `!(${lCheck} && ${rCheck}) ? undefined : (${lValue} ${op} ${rValue})`
    }
  })

while (Object.keys(unsolvedWires).length > 0) {
  const latestUnsolvedWires = Object.keys(unsolvedWires)
    .reduce((u, w) => {
      // eslint-disable-next-line no-eval
      const evalResult = eval(unsolvedWires[w])
      if (evalResult === undefined) {
        u[w] = unsolvedWires[w]
      } else {
        solvedWires[w] = evalResult
      }
      return u
    }, {})

  unsolvedWires = latestUnsolvedWires
}

console.log(solvedWires.a)
