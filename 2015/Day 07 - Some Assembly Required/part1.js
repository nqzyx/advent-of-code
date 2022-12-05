const { readFileSync } = require('fs')

var circuit = readFileSync('input.txt', 'utf8')
  .split('\n')
  .map((spec) => {
    const [input, wire] = spec.split(' -> ')
    if (input.match(/^\d+$/)) {
      // Input: Value
      return {
        type: 'wire',
        name: wire,
        input: { type: 'value', value: parseInt(input) },
      }
    } else if (input.match(/^[a-z]+$/)) {
      // Input: Wire
      return {
        type: 'wire',
        name: wire,
        input: { type: 'wire', name: input },
      }
    } else if (input.match(/^NOT ([a-z]+)$/)) {
      // Input: NOT Gate
      const [op, wire] = input.split(' ')
      return {
        type: 'wire',
        name: wire,
        input: {
          type: 'gate',
          name: op.toLowerCase(),
          input: { type: 'wire', name: wire },
        },
      }
    } else {
      // Input: Gate
      let [left, op, right] = input.split(' ')
      left = left.match(/^\d+$/)
        ? { type: 'value', value: parseInt(left) }
        : left.match(/^[a-z]+$/)
        ? { type: 'wire', name: left }
        : `Unknown left input: ${left}`
      right = right.match(/^\d+$/)
        ? { type: 'value', value: parseInt(right) }
        : right.match(/^[a-z]+$/)
        ? { type: 'wire', name: right }
        : `Unknown right input: ${right}`
      return {
        type: 'wire',
        name: wire,
        input: {
          type: 'gate',
          name: op.toLowerCase(),
          input: { type: 'pair', left, right },
        },
      }
    }
  })

console.log(circuit)

let wireValueUpdated = false
do {
  newCircuit = circuit.map((wire) => {
    if (Object.keys(wire).includes('value') && wire.value) {
      return wire;
    }
    if (wire.input.type === 'value') {
      wire.value = wire.input.value
      wireValueUpdated = true
    } else if (wire.input.type === 'wire') {
      const inputWire = circuit.find((w) => w.name === wire.input.name)
      if (inputWire?.value) {
        wire.value = inputWire.value
        wireValueUpdated = true
      }
    } else if (wire.input.type === 'gate' && wire.input.name === 'not') {
      const inputWire = circuit.find((w) => w.name === wire.input.input.name)
      if (inputWire?.value) {
        wire.value = ~inputWire.value & 0xffff
        wireValueUpdated = true
      }
    } else if (wire.input.type === 'gate') {
      const { name, input } = wire.input
      if (input.type === 'pair') {
        const { left, right } = input
        if (left.type === 'wire' && right.type === 'wire') {
          const leftWire = circuit.find((w) => w.name === left.name)
          const rightWire = circuit.find((w) => w.name === right.name)
          if (leftWire?.value && rightWire?.value) {
            wire.value =
              name === 'and'
                ? leftWire.value & rightWire.value
                : name === 'or'
                ? leftWire.value | rightWire.value
                : name === 'lshift'
                ? leftWire.value << rightWire.value
                : name === 'rshift'
                ? leftWire.value >> rightWire.value
                : `Unknown gate: ${name}`
            wireValueUpdated = true
          }
        }
      } else {
        wire.value = `Unknown input: ${wire.input}`
      }
    }
    return wire;
  })
  
  if (wireValueUpdated) {
    circuit = newCircuit
  }
} while (wireValueUpdated)

console.log(circuit);
