// AoC 2022 Day 16 - Proboscidea Volcanium - Part 1

import { readFileSync } from 'fs'
import dijkstrajs from 'dijkstrajs'

import { create, all, forEach } from 'mathjs'

const config = {}
const math = create(all, config)

const origin = 'AA'
const timeLimit = 30
const moveTime = 1
const valveTime = 1
let maxCost = 0

function valve ({name, flowRate, isOpen, neighbors}) {
  this.name = name
  this.isOpen = isOpen
  this.flowRate = flowRate instanceof Number 
    ? flowRate 
    : parseInt(flowRate.toString())
  this.currentFlow = () => {
    return (this.isOpen) ? this.flowRate : 0
  }
  this.neighbors = neighbors instanceof Array
    ? neighbors
    : neighbors.toString().split(',').map((name) => name.trim())
  this.cost = () => {
    return (this.currentFlow() === 0)
      ? maxCost
      : maxCost - this.flowRate
  }  
}

function path ({}) {
  
}

const valves = readFileSync('input.txt', 'utf8')
  .split('\n')
  // extract data from the spec
  .map((spec, specNdx, specArr) => {
    const valveRegex =
      /^Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)$/
    const [, name, flowRate, neighbors] = spec.match(valveRegex)
    maxCost = math.max(maxCost, flowRate)
    return new valve({ 
      name, 
      flowRate, 
      isOpen: false,
      neighbors
    })
  })

// console.log('\n=====', 'valveSpecs', '=====')
// console.log(valves)
// process.exit(0)

const getValveMap = (valves) =>
  valves.reduce((vMap, valve, vi, vArr) => {
    const neighbors = valve.neighbors.reduce((nMap, neighborName, ni, nArr) => {
      nMap[neighborName] = vArr.find((v) => v.name === neighborName).cost()
      return nMap
    }
    , {})
    vMap[valve.name] = neighbors
    return vMap
  }, {})

// console.log('\n=====', 'valveMap', '=====')
// console.log(getValveMap(valves))
// process.exit(0)

const adjustPath = (path) => {
      let minimumTotalFlow
      let { timeRemaining, cumulativeFlow, currentFlow } = path
      return Object.assign(path, {
        route: path.route.reduce((newRoute, v, vi, vArr) => {
          const vTime =
            (vi > 0 ? 1 : 0) + //move time
            (v.ratedFlow > 0 ? 1 : 0) // open time
          if (timeRemaining > vTime) {
            v.isOpen = true
            ;(timeRemaining -= vTime),
              (cumulativeFlow += currentFlow * vTime),
              (currentFlow += v.currentFlow())
            minimumTotalFlow = cumulativeFlow + currentFlow * timeRemaining
            newRoute.push(v)
          }
          return newRoute
        }, []),
        timeRemaining,
        cumulativeFlow,
        currentFlow,
        minimumTotalFlow,
      })
}

const getPaths = (start, valves, timeRemaining = timeLimit) => {
  return valves
    .filter((valve, vi, vArr) => !(valve.name == start || valve.currentFlow() === 0))
    .map((valve, vi, vArr) => {
      let newRoute
      try {
        newRoute = dijkstrajs
          .find_path(getValveMap(valves), start, valve.name)
          .map((vName) => valves.find((vSpec) => vSpec.name === vName))
      } catch {
        newRoute = []
      }
      return {
        from: start,
        to: valve.name,
        timeRemaining: timeRemaining,
        currentFlow: 0,
        cumulativeFlow: 0,
        route: newRoute
      }
    })
    .map((path, pi, pArr) => {
      adjustPath(path)
    })
  }
  
// console.log('\n====', `All possible paths from node '${origin}'`, '====')
// const xPaths = getPaths(origin, valves)
// xPaths.forEach((path, pi) => {
//   console.log(pi, ':', path)
// })
// console.log('\n====', 'Best Candidate Path', '====')
// console.log(xPaths.sort((pathA, pathB) => pathB.minimumTotalFlow - pathA.minimumTotalFlow)[0])
// process.exit(0)

const paths = getPaths(origin, valves)
    .sort((pathA, pathB) => pathB.minimumTotalFlow - pathA.minimumTotalFlow)
    .map((path, pi, pArr) => {
      const route = path.route
      const unvisitedValves = valves
        .filter((v) => )
        .filter((v) => !(route.map((n) => n.name)).includes(v.name))
      const extensionPaths = getPaths(route[route.length - 1].name, unvisitedValves, path.timeRemaining)
      if (extensionPaths.shift() !== undefined) {
        extensionPaths.reduce((x, path, pi, pArr) => {

        }, null)
      }
    })

console.log('\n====', 'bestPath', '====')
console.log(bestPath)
process.exit(0)
