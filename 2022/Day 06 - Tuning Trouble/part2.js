const { readFileSync } = require('fs')

const buffer = readFileSync('input.txt', 'utf8').split('')

const pml = 4
const mml = 14

for (let c = 0; c < buffer.length - pml; c++) {
  const code = buffer.slice(c, c + pml)
  const cuc = new Set(code).size
  if (cuc < pml) continue
  for (let m = c + pml; m < buffer.length; m++) {
    const msg = buffer.slice(m, m + mml)
    const muc = new Set(msg).size
    if (muc < mml) {
      continue
    }
    console.log(m + mml)
    break
  }
  break
}
