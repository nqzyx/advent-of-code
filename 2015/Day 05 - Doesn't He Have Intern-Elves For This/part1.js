const { readFileSync } = require('fs')

console.log(
  readFileSync('input.txt', 'utf8')
    .split('\n')
    .reduce(
      (nice, s) =>
        (nice += s.match(/(ab|cd|pq|xy)/) ? 0 : s.match(/([aeiou].*){3,}/) && s.match(/(.)\1/) ? 1 : 0),
      0
    )
)
