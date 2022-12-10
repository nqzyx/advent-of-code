// AoC 2022 Day 08 - xxx - Part 1

// const { readFileSync } = require('fs')
import { readFileSync } from 'fs'

import { create, all } from 'mathjs'
const config = {}
const math = create(all, config)

const forest = math.matrix(readFileSync('input.txt', 'utf8').split('\n').map((r) => r.split('')))

const [colLength, rowLength] = forest.size()
const [maxCol, maxRow] = math.subtract([colLength, rowLength], 1)
const totalTrees = colLength * rowLength

const treeVisibility = forest.map(
  (height, [tCol, tRow], f) =>
    // Is the tree on perimeter?
    [0, maxCol].includes(tCol) ||
    [0, maxRow].includes(tRow) ||
    // Are all the trees to the north shorter?
    height > math.max(math.subset(f, math.index(tCol, math.range(0, tRow)))) ||
    // Are all the trees to the south shorter?
    height > math.max(math.subset(f, math.index(tCol, math.range(tRow + 1, maxRow, true)))) ||
    // Are all the trees to the east shorter?
    height > math.max(math.subset(f, math.index(math.range(tCol + 1, maxCol, true), tRow))) ||
    // Are all the trees to the west shorter?
    height > math.max(math.subset(f, math.index(math.range(0, tCol), tRow)))
)

const visibleTrees = math.count(math.filter(math.flatten(treeVisibility), (v) => v))

console.log(visibleTrees, 'of', totalTrees, 'trees are visible')
