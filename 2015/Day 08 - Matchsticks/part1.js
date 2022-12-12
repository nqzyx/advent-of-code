// AoC 2015 Day 08 - Matchsticks - Part 1

const { readFileSync } = require('fs')

const escChar = '\\'
const dQuote = '"'
const lcX = 'x'

console.log(
  readFileSync('input.txt', 'utf8')
    .split('\n')
    .map((line) => {
      console.log('Line', line, '\nlength', line.length)
      const ol = line.length // original string length
      let [, s] = line.match(/^"(.*)"$/)
      s.replace(/""/, '')
      return (() => {
        let ps, pe, rs, sl, sr
        while ((ps = s.indexOf(escChar)) !== -1) {
          sl = s.substring(0, ps + 1)
          switch (s[ps + 1]) {
            case dQuote:
              pe = ps + 2
              rs = dQuote
              break
            case lcX:
              pe = ps + 4
              rs = String.fromCharCode(parseInt(s.substring(ps, pe), 16))
              break
            default:
              pe = ps
              rs = ''
              break
          }
          sr = pe < s.length ? s.substring(pe) : ''
          s = sl + rs + sr
        }
        const retVal = { ol, s, cl: s.length, diff: ol - s.length }
        console.log(JSON.stringify(retVal))
        return retVal
      })()
    })
)
