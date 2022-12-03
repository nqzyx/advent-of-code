const { readFileSync } = require('fs')

console.log(
  readFileSync('./input.txt', 'utf-8')
    .split('\n')
    .map(pkg =>
      pkg
        .split('x')
        .map(Number)
        .sort((a, b) => Number(a) - Number(b))
    )
  .reduce((acc, [l, w, h]) => acc + (2*l + 2*w +  + l*w*h), 0)
)
