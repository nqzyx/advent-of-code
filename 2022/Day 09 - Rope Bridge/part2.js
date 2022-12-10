// AoC 2022 Day 09 - Rope Bridge - Part 2

import { readFileSync } from 'fs'

import { create, all } from 'mathjs'
const config = {}
const math = create(all, config)

const instructions = readFileSync('input.txt', 'utf8').split('\n').map((r) => { const [dir, distance] = r.split(' '); return [dir, parseInt(distance)]})

const s = { x: 0, y: 0 }
const knots = math.range(0, 10).map((i) =>({x:0, y:0})).toArray()


function follow (leader, follower) { 
    if (math.range(leader.x - 1, leader.x + 1, true).toArray().includes(follower.x) &&
        math.range(leader.y - 1, leader.y + 1, true).toArray().includes(follower.y)) {
            return follower
        }
    
    follower.x += (leader.x === follower.x) ? 0 : (leader.x > follower.x) ? 1 : -1
    follower.y += (leader.y === follower.y) ? 0 : (leader.y > follower.y) ? 1 : -1
    return follower
}

console.log(
    math.size(
        instructions.reduce((visited, instruction) => {
            const [dir, reps] = instruction
            for (let r = 0; r < parseInt(reps); r++) {
                switch (dir) {
                    case 'U':
                        knots[0].y += 1
                        break

                    case 'D':
                        knots[0].y -= 1
                        break

                    case 'L':
                        knots[0].x -= 1
                        break
                    
                    case 'R':
                        knots[0].x += 1
                        break
                }
                knots.reduce((newKnots, knot, ki, oldKnots) => {
                    if (ki === 0) {
                        newKnots.push(knot)
                    } else if (ki > 0) {
                        newKnots.push(Object.assign({}, follow(oldKnots[ki - 1], knot)))
                    }
                    return newKnots
                }, [])
                const tk = knots[knots.length - 1]
                if (visited.find((v) => v.x === tk.x && v.y === tk.y) === undefined) {
                    visited.push({...tk})
                }
            }
            return visited
        }, [s])
    )[0]
)
