const { readFileSync } = require('fs')

const stacks = []

const [initialStacks, moves] = readFileSync('input.txt', 'utf8').split(
  '\n\n'
)

initialStacks
  .split('\n')
  .reverse()
  .map((level) => {
    for (let s = 0; s < (level.length + 1) / 4; s++) {
      const crate = level.slice(s * 4, s * 4 + 4)
      if (crate.match(/^\[[A-Z]\]/)) {
        if (!(stacks[s] instanceof Array)) stacks[s] = []
        stacks[s].push(crate.slice(1, 2))
      }
    }
    return level
  })

moves
  .split('\n')
  .map((move) => {
    const [ _, boxes, from, to ] = move.match(/^move (\d+) from (\d+) to (\d+)$/)
    stacks[to - 1].push(...stacks[from - 1].splice(-boxes))
    return move
  })

console.log(stacks.map((stack) => stack.slice(-1)[0]).join(''))
