const { readFileSync } = require('fs')

function sectionsOverlap([start1, end1], [start2, end2]) {
  return (
    (start2 >= start1 && start2 <= end1) ||
    (end2 >= start1 && end2 <= end1) ||
    (start1 >= start2 && start1 <= end2) ||
    (end1 >= start2 && end1 <= end2)
  )
}

console.log(
  readFileSync('./input.txt', 'utf-8')
    .split('\n')
    .map((assignments) =>
      assignments.split(',').map((sections) => sections.split('-').map(Number))
    )
    .reduce(
      (overlaps, assignments) =>
        overlaps + (sectionsOverlap(assignments[0], assignments[1]) ? 1 : 0),
      0
    )
)
