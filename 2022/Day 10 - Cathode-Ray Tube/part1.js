// AoC 2022 Day 10 - Cathode-ray Tube - Part 1

import { readFileSync } from 'fs'

import { create, all } from 'mathjs'
const config = {}
const math = create(all, config)

const program = readFileSync('input.txt', 'utf8').split('\n')
const initialState = { xValue: 1, addend: 0 }

let cycles = program.reduce((c, statement) => {
  // console.log(`c: ${JSON.stringify(c)}`)
  const lastCycleState = c.length > 0 ? c.slice(-1)[0] : initialState
  const newXValue = lastCycleState.xValue + lastCycleState.addend
  const [, instruction, addend] = statement.match(/^(noop|addx) ?(-?\d+)?$/)
  switch (instruction) {
    case 'noop':
      c.push({ statement, xValue: newXValue, addend: 0 })
      break
    case 'addx':
      c.push({ statement, xValue: newXValue, addend: 0 })
      c.push({ statement, xValue: newXValue, addend: parseInt(addend) })
      break
    default:
      throw new Error(`Invalid instruction (${instruction})`)
  }
  return c
}, [])

cycles = cycles.map((c, ci) => {
  const c2 = Object.assign(c, { signalStrength: (ci + 1) * c.xValue })
  // console.log(ci, c2)
  return c2
})

const indices = math.range(0, 6).map((i) => 19 + i * 40)

const values = math.subset(cycles, math.index(indices))

console.log(values.reduce((s, c) => s + c.signalStrength, 0))
