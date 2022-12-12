// AoC 2022 Day 13 - xxx - Part 2

import { readFileSync } from 'fs'

import dijkstrajs from 'dijkstrajs'

import { create, all } from 'mathjs'
import { json } from 'stream/consumers'

const input = readFileSync('input.txt', 'utf8').split('\n')

console.log(JSON.stringify(input))
