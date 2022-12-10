// AoC 2022 Day 10 - Cathode-ray Tube - Part 2

import { readFileSync } from 'fs'

import { create, all } from 'mathjs'
const config = {}
const math = create(all, config)

const program = readFileSync('input.txt', 'utf8').split('\n')
const pixelCount = 240
const initialState = { sequence: 0, statement: null, xValue: 1, addend: 0, pixel: null }

let cycles = program.reduce((c, statement) => {
  const lastCycleState = c.slice(-1)[0] ?? initialState
  const xValue = lastCycleState.xValue + lastCycleState.addend
  let sequence = lastCycleState.sequence + 1
  let pixel = (sequence % pixelCount) - 1
  let [ , instruction, addend] = statement.match(/^(noop|addx) ?(-?\d+)?$/)

  switch (instruction) {
    case 'noop' :
      c.push({ sequence, xValue, statement, addend: 0, pixel })
      break
    case 'addx' :
      c.push({ sequence, xValue, statement, addend: 0, pixel })
      sequence++
      addend = parseInt(addend)
      pixel++
      c.push({ sequence, xValue, statement, addend, pixel })
      break
    default:
      throw new Error(`Invalid instruction (${instruction})`)
  }
  return c
}, [])

const pixels = cycles.reduce((pixels, c) => {
  if (math.range(c.xValue - 1, c.xValue + 1, true).toArray().includes(c.pixel % 40)) {
    pixels[c.pixel] = 'X'
  }
  return pixels
}, Array(pixelCount).fill(' '))

math.reshape(pixels, [6,40]).forEach((pr) => { console.log(pr.join(''))})
