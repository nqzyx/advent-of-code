const { readFileSync } = require('fs')

console.log(
  readFileSync('input.txt', 'utf8')
    .split('\n')
    .reduce(
      (nice, s) =>
        nice + (
          s.match(/(..).*\1/) &&
          s.match(/(.).\1/)
            ? 1
            : 0
        ),
      0
    )
)
