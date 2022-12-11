// AoC 2022 Day 12 - xxx - Part 1

import { readFileSync } from 'fs'

import { create, all, min } from 'mathjs'
import { parse } from 'path'
const config = {}
const math = create(all, config)

const input = readFileSync('input.txt', 'utf8').split('\n\n')
