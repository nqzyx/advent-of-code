const { readFileSync } = require('fs')

const buffer = readFileSync('input.txt', 'utf8').split('')

const l = 4

for (let c = 0; c < buffer.length - l; c++) {
  const code = buffer.slice(c, c + l)
  const uc = new Set(code).size
  if (uc < l) continue
  console.log(c + l)
  break
}
