// AoC 2022 Day 09 - Rope Bridge - Part 1

import { readFileSync } from 'fs'

import { create, all } from 'mathjs'
const config = {}
const math = create(all, config)

const instructions = readFileSync('input.txt', 'utf8').split('\n').map((r) => { const [dir, distance] = r.split(' '); return [dir, parseInt(distance)]})

const h = { x: 0, y: 0 }
const s = { x: 0, y: 0 }
const t = { x: 0, y: 0 }

function follow (h, t) { 
    if (math.range(h.x - 1, h.x + 1, true).toArray().includes(t.x) &&
        math.range(h.y - 1, h.y + 1, true).toArray().includes(t.y)) {
            return t
        }
    
    t.x += (h.x === t.x) ? 0 : (h.x > t.x) ? 1 : -1
    t.y += (h.y === t.y) ? 0 : (h.y > t.y) ? 1 : -1
    return t
}

console.log(
    math.size(
        instructions.reduce((visited, instruction) => {
            const [dir, reps] = instruction
            for (let r = 0; r < parseInt(reps); r++) {
                switch (dir) {
                    case 'U':
                        h.y += 1
                        break

                    case 'D':
                        h.y -= 1
                        break

                    case 'L':
                        h.x -= 1
                        break
                    
                    case 'R':
                        h.x += 1
                        break
                }
                Object.assign(t, follow({...h}, {...t}))
                if (visited.find((v) => v.x === t.x && v.y === t.y) === undefined) {
                    visited.push({...t})
                }
            }
            return visited
        }, [s])
    )[0]
)
