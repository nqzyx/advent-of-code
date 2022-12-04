const { readFileSync } = require('fs')

console.log(
  readFileSync('./input.txt', 'utf-8')
    .split('\n')
    .map((assignments) =>
      assignments.split(',').map((sections) => sections.split('-').map(Number))
    )
    .reduce((duplicates, assignments) => {
      const [start1, end1] = assignments[0]
      const [start2, end2] = assignments[1]
      if (
        (start1 <= start2 && end1 >= end2) ||
        (start2 <= start1 && end2 >= end1)
      ) {
        duplicates++
      }
      return duplicates
    }, 0)
)
