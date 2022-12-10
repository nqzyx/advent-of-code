// AoC 2022 Day 08 - Treetop Tree House - Part 2

/* eslint-disable operator-linebreak */

import { readFileSync } from 'fs'
import { create, all } from 'mathjs'
const config = {}
const math = create(all, config)

const forest =
  readFileSync('input.txt', 'utf8')
    .split('\n')
    .map((r) => r.split('').map((h) => parseInt(h)))

const [maxColumn, maxRow] = math.subtract(math.size(forest), 1)

function getScenicScoreForTree (thisTreeHeight, row, column, forest) {
  // By definition, trees on the edge have a scenicScore of 0
  if (row === 0 || row === maxRow || column === 0 || column === maxColumn) {
    return 0
  }

  const treeRow = math.flatten(math.row(forest, row))
  const treeColumn = math.flatten(math.column(forest, column))
  const trees = {
    north: treeColumn.slice(0, row).reverse(),
    south: treeColumn.slice(row + 1, maxRow + 1),
    east: treeRow.slice(column + 1, maxColumn + 1),
    west: treeRow.slice(0, column).reverse()
  }

  try {
    /* eslint-disable multiline-ternary */
    /* eslint-disable indent */
    const northScenicScore =
      trees.north.length === 0
        ? 0
        : getDirectionalScenicScore(thisTreeHeight, trees.north)
    const southScenicScore =
      trees.south.length === 0
        ? 0
        : getDirectionalScenicScore(thisTreeHeight, trees.south)
    const eastScenicScore =
      trees.east.length === 0
        ? 0
        : getDirectionalScenicScore(thisTreeHeight, trees.east)
    const westScenicScore =
      trees.west.length === 0
        ? 0
        : getDirectionalScenicScore(thisTreeHeight, trees.west)
    return (
      northScenicScore * southScenicScore * eastScenicScore * westScenicScore
    )
    /* eslint-enable indent */
    /* eslint-enable multiline-ternary */
  } catch (error) {
    throw new Error(
      JSON.stringify(
        {
          thisTreeHeight,
          column,
          maxCol: maxColumn,
          row,
          maxRow,
          trees: {
            north: JSON.stringify(trees.north),
            south: JSON.stringify(trees.south),
            east: JSON.stringify(trees.east),
            west: JSON.stringify(trees.west)
          },
          innerError: error.message
        },
        null,
        2
      ) + '\n'
    )
  }

  function getTreeInfo (thisTreeHeight, thatTreeHeight, visible) {
    const treeInfo = {}
    treeInfo.scenicScore = visible ? 1 : 0
    treeInfo.nextTreeVisible = visible && thatTreeHeight < thisTreeHeight
    return treeInfo
  }

  function getDirectionalScenicScore (thisTreeHeight, trees) {
    let nextTreeVisible = true
    const treesScore = trees.reduce((treesScore, thatTreeHeight) => {
      const treeInfo = getTreeInfo(
        thisTreeHeight,
        thatTreeHeight,
        nextTreeVisible
      )
      nextTreeVisible = treeInfo.nextTreeVisible
      return treesScore + treeInfo.scenicScore
    }, 0)
    return treesScore
  }
}

const scenicScores = forest.map((rowOfTrees, row, forest) =>
  rowOfTrees.map((treeHeight, column) =>
    getScenicScoreForTree(treeHeight, row, column, forest)
  )
)

console.log(math.max(scenicScores))
