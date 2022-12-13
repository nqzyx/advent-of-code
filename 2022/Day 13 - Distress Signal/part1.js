// AoC 2022 Day 13 - Distress Signal - Part 1

import { readFileSync } from 'fs'

const compare = (left, right) => {
  console.log('- Compare', left, 'vs', right)
  console.group()
  let result
  if (Number.isInteger(left) && Number.isInteger(right)) {
    result = left - right
    if (result < 0) {
      console.log('- Left side is smaller, so inputs ARE in the right order')
    } else if (result > 0) {
      console.log(
        '- Right side is smaller, so inputs ARE NOT in the right order'
      )
    }
  } else if (Array.isArray(left) && Number.isInteger(right)) {
    console.log('- Mixed types: convert right to', right, 'and retry comparison')
    result = compare(left, [right])
  } else if (Array.isArray(right) && Number.isInteger(left)) {
    console.log('- Mixed types: convert left to', left, 'and retry comparison')
    result = compare([left], right)
  } else if (Array.isArray(left) && Array.isArray(right)) {
    let i = 0
    result = 0
    while (i < Math.min(left.length, right.length) && result === 0) {
      result = compare(left[i], right[i])
      if (result !== 0) break
      result = 0
      i += 1
    }
    if (result === 0) {
      if (left.length < right.length) {
        console.log('- Left side ran out of items, so inputs ARE in the right order')
        result = -1
      } else if (left.length > right.length) {
        console.log('- Right side ran out of items, so input ARE NOT in the right order')
        result = 1
      }
    }
  } 
  console.groupEnd()
  return result
}

const input = readFileSync('input.txt', 'utf8')
  .split('\n\n')
  .reduce((sum, p, pi) => {
    console.log(`== Pair ${pi + 1} ==`)
    const [lStr, rStr] = p.split('\n')
    const compareResult = compare(eval(lStr), eval(rStr))
    sum += (compareResult < 0) ? pi + 1 : 0
    return sum
  }, 0)

console.log('\nAnswer:', JSON.stringify(input))
