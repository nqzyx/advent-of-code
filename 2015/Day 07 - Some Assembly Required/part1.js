const { readFileSync } = require('fs')

let wires = readFileSync('input.txt', 'utf8')
  .split('\n')
  .reduce((updatedWires, connection) => {
    const [input, wire] = connection.split(' -> ')
    if (input.match(/^\d+$/)) {
      // Input is a constant
      const value = parseInt(input)
      updatedWires[wire] = { input: { type: 'value', value }, output: value }
    } else if (input.match(/^[a-z]+$/)) {
      // Input is a wire
      updatedWires[wire] = { input: { type: 'wire', name: input }, output: null }
    } else if (input.match(/^NOT ([a-z]+)$/)) {
      // Input is a NOT gate
      const [op, inputWire] = input.split(' ')
      updatedWires[wire] = { input: { type: 'gate', name: op.toLowerCase(), inputs: [{ type: 'wire', name: inputWire }] }, output: null }
    } else if (input.match(/^([a-z]+) (AND|OR|LSHIFT|RSHIFT) ([a-z]+)$/)) {
      // Input is a gate
      let [left, op, right] = input.split(' ')
      op = op.toLowerCase()
      left = left.match(/^\d+$/) ? parseInt(left) : { type: 'wire', name: left }
      right = right.match(/^\d+$/) ? parseInt(right) : { type: 'wire', name: right }
      updatedWires[wire] = { input: { type: 'gate', name: op, inputs: [left, right] }, output: null }
    }
    return updatedWires
  }, {})

function getWireValue (name) {
  const wire = wires[name]
  return wire?.output ?? null
}

let valuesUpdated = false
do {
  wires = Object.keys(wires).reduce((updatedWires, wireName) => {
    valuesUpdated = false
    const wire = wires[wireName]
    const input = wire.input

    console.log(`${wireName}: ${JSON.stringify(wire)}`)

    switch (input.type) {
      case 'value':
        break

      case 'wire':
        // eslint-disable-next-line no-case-declarations
        const value = getWireValue(input.name)
        if (value instanceof Number) {
          Object.assign(wire, { input: { type: 'value', value }, output: value })
          valuesUpdated = true
        }
        break

      case 'gate':
        switch (input.name) {
          case 'not':
            // eslint-disable-next-line no-case-declarations
            const value = getWireValue(input.inputs[0].name)
            if (value instanceof Number) {
              Object.assign(wire, { input: { type: 'value', value: ~value & 0xffff }, output: ~value & 0xffff })
              valuesUpdated = true
            }
            break

          default: // and || or || lshift || rshift
            input.inputs.forEach((inputWire) => {
              const inputWireValue = getWireValue(inputWire.name)
              if (inputWireValue instanceof Number) {
                Object.assign(inputWire, { input: { type: 'value', value: inputWireValue }, output: inputWireValue })
                valuesUpdated = true
              }
            })
            if (input.inputs.every((inputWire) => inputWire.output instanceof Number)) {
              const [left, right] = input.inputs
              const value = input.name === 'and'
                ? left.output & right.output
                : input.name === 'or'
                  ? left.output | right.output
                  : input.name === 'lshift'
                    ? left.output << right.output
                    : left.output >> right.output
              Object.assign(wire, { input: { type: 'value', value }, output: value })
              valuesUpdated = true
            }
            break
        }
        break
    }
    updatedWires[wireName] = wire
    return updatedWires
  }, {})
} while (valuesUpdated)

console.log(`Wires length: ${wires.length}`)
