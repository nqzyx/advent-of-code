// AoC 2022 Day 15 - Beacon Exclusion Zone - Part 1

import { readFileSync } from 'fs'
import lodash from 'lodash'
const { isEqual } = lodash

import { create, all, zeros } from 'mathjs'
const config = {}
const math = create(all, config)

const searchLimit = 4000000

const start = {
  min: { x: Infinity, y: Infinity },
  max: { x: -Infinity, y: -Infinity },
}

let { x: minX, y: minY } = start.min
let { x: maxX, y: maxY } = start.max

const getDist = ([aX, aY], [bX, bY]) => math.abs(aX - bX) + math.abs(aY - bY)
const getCorners = ([oX, oY], d) => [
  [oX, oY + d],
  [oX, oY - d],
  [oX + d, oY],
  [oX - d, oY],
]

const sensors = readFileSync('input.txt', 'utf8')
  .split('\n')
  .map((s) => {
    const [, sX, sY, bX, bY] = s.match(
      /^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$/
    )
    const pos = [parseInt(sX), parseInt(sY)]
    const b = [parseInt(bX), parseInt(bY)]
    const d = getDist(pos, b)
    const c = getCorners(pos, d)
    const [maxX, maxY] = math.max([pos, b, ...c], 0)
    const [minX, minY] = math.min([pos, b, ...c], 0)
    return { pos, b, d, minX, minY, maxX, maxY }
  })

sensors.forEach((r) => {
  minX = math.min(minX, r.minX)
  maxX = math.max(maxX, r.maxX)
  minY = math.min(minY, r.minY)
  maxY = math.max(maxY, r.maxY)
})

console.log({ minX, maxX, minY, maxY })

const xRange = math.range(0, 20, true)
const yRange = math.range(0, 20, true)

let beaconLocation = null

for (let x = 0; x <= searchLimit; x++) {
    for (let y = 0; y <= searchLimit; y++) {
        const loc = [x, y]
        if (sensors.some((s) => !(s.d < getDist(s.pos, loc)))) continue
        beaconLocation = loc
        break
    }
    if (beaconLocation !== null) break
}

console.log(beaconLocation)

console.log((beaconLocation[0] * 4000000) + beaconLocation[1])
