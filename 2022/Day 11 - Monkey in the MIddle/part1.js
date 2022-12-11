// AoC 2022 Day 11 - Monkey in the Middle - Part 1

import { readFileSync } from 'fs'

import { create, all, min } from 'mathjs'
import { parse } from 'path'
const config = {}
const math = create(all, config)

const numRounds = 20
const worryMitigator = 3;

const monkeySubset = () => ['name', 'itemsInspected']

const monkeyInputs = readFileSync('input.txt', 'utf8').split('\n\n')

const makeInspectItemMethod = (opSpec) => {
  const op = opSpec.replace(/old/g,'i')
  return function inspectItem(i) {
    return math.floor(eval(op) / worryMitigator)
  }
}

const makeItemsList = (list) => list.split(',').map((i) => parseInt(i.trim()))

const makeTakeTurnMethod = () =>
function takeTurn() {
  this.items.forEach((item) => {
    item = this.inspectItem(item)
    this.itemsInspected++
    this.throwItem(item)
  })
  this.items.length = 0
}

const makeThrowItemMethod = (d, trueMonkey, falseMonkey) => {
  const divisor = parseInt(d)
  const trueMonkeyIndex = parseInt(trueMonkey)
  const falseMonkeyIndex = parseInt(falseMonkey)
  return function throwItem(i) {
    i % divisor === 0
    ? monkeys[trueMonkeyIndex].items.push(i)
    : monkeys[falseMonkeyIndex].items.push(i)
  }
}

let monkeys = monkeyInputs.map((mInput, mIndex) => {
  const m = { }
  const mLines = mInput.split('\n').map((l) => l.toLowerCase().trim())
  mLines.forEach((ml, mli, mla) => {
    const cmd = ml.match(/^([A-Za-z]+)/)[1]
    switch (cmd) {
      case 'monkey':
        const mi = parseInt(ml.match(/^monkey (\d+):$/)[1])
        if (mi !== mIndex) {
          throw new Error(`Invalid Monkey number: Expected ${mIndex}, Actual ${mi}`)
        }
        m.name = ml.slice(0, ml.length - 1)
        break

      case 'starting':
        m.items = makeItemsList(ml.match(/^starting items: (.*)$/)[1])
        break

      case 'operation':
        m.inspectItem = makeInspectItemMethod(ml.match(/^operation: new = (.*)$/)[1])
        break

      case 'test':
        const d = ml.match(/^test: divisible by (\d+)$/)[1]
        const mTrue = mla[mli + 1].match(/^if true: throw to monkey (\d+)$/)[1]
        const mFalse = mla[mli + 2].match(/^if false: throw to monkey (\d+)$/)[1]
        m.throwItem = makeThrowItemMethod(d, mTrue, mFalse)
        break
    }
  })
  m.takeTurn = makeTakeTurnMethod()
  m.itemsInspected = 0
  return m
})

const rounds = math.range(0, numRounds).toArray()

const selectedRounds = rounds.reduce(
  (s, r) => (r === 0 || r === 19 || (r + 1) % 1000 === 0 ? [...s, r] : s),
  []
)

rounds.forEach((round) => {
  monkeys.forEach((m) => {
    m.takeTurn()
  })
  if (selectedRounds.includes(round)) {
    console.log(`== After round ${round + 1} ==`)
    monkeys.forEach((m, mi) => {
      console.log(
        `Monkey ${mi} inspected items ${
          m.itemsInspected
        } times. ${JSON.stringify(m, monkeySubset())}`
      )
    })
  }
})

// console.log(('=== Top Monkeys ' + '='.repeat(80)).slice(0, 80))

const topMonkeys = monkeys.sort((m1, m2) => m2.itemsInspected - m1.itemsInspected)
  
topMonkeys.forEach((m, mi) => {
    // console.log(`Monkey[${mi}]: ${JSON.stringify(m)}`)
  })

console.log(`\nTop Monkeys:\n\n${JSON.stringify(monkeys, [...monkeySubset(), 'items'])}`)

console.log(
  topMonkeys.slice(0,2)
    .reduce((total, m) =>
      total * m.itemsInspected
      , 1
    )
)
