// AoC 2022 Day 12 - Hill Climbing Algorithm - Part 1

import { readFileSync } from 'fs'

import dijkstrajs from 'dijkstrajs'

import { create, all } from 'mathjs'

const config = {}
const math = create(all, config)

let origin = [0, 0]
let destination = [0, 0]

const getHeight = (h) => h.charCodeAt(0) - 'a'.charCodeAt(0)
const getName = ([row, column]) => `r${row}c${column}`

const heightMap = readFileSync('input.txt', 'utf8')
  .split('\n')
  .map((row, ri) =>
    row.split('').map((column, ci) => {
      let h
      switch (column) {
        case 'S':
          origin = [ri, ci]
          h = getHeight('a')
          break

        case 'E':
          destination = [ri, ci]
          h = getHeight('z')
          break

        default:
          h = getHeight(column)
          break
      }
      return {
        name: getName([ri, ci]),
        height: h,
      }
    })
  )

const [maxRow, maxColumn] = math.subtract(math.size(heightMap), 1)

const getDifficulty = (h1, h2) => {
  let d = h2 - h1
  return d < 0 ? 0 : d === 0 ? 1 : d === 1 ? 2 : -1
}

const getNeighbor = (h, map, row, col) => {
  const n = map[row][col]
  const e = {}
  const d = getDifficulty(h, n.height)
  if (n && d >= 0) {
    e[n.name] = d
  }
  return e
}

const graph = heightMap.reduce(
  (graph, row, ri, map) =>
    Object.assign(
      graph,
      row.reduce((graph, col, ci) => {
        const edges = {}

        if (ri > 0)
          Object.assign(edges, getNeighbor(col.height, map, ri - 1, ci))
        if (ri < maxRow)
          Object.assign(edges, getNeighbor(col.height, map, ri + 1, ci))
        if (ci > 0)
          Object.assign(edges, getNeighbor(col.height, map, ri, ci - 1))
        if (ci < maxColumn)
          Object.assign(edges, getNeighbor(col.height, map, ri, ci + 1))

        graph[col.name] = edges

        return graph
      }, {})
    ),
  {}
)

console.log(
  dijkstrajs.find_path(graph, getName(origin), getName(destination)).length - 1
)
