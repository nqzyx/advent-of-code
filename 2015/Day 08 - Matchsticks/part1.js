// AoC 2015 Day 08 - Matchsticks - Part 1

const { readFileSync } = require('fs')

const initialInput =
  readFileSync('input.txt', 'utf8')
    .split('\n')

const unquotedInput =
  initialInput
    .map((ol, li, f) =>
      // Remove leading and trailing quotes
      ol.replace(/^"(.*)"$/, '$1')
    )

const backSlashedInput =
  unquotedInput
    .map((ol, li, f) =>
      // Escaped backslash (\\)
      ol.replace(/\\\\/g, '\\')
    )

const deHexedInput =
  backSlashedInput
    .map((ol, li, f) =>
      // Escaped hex char code (\x20)
      ol.replace(/\\x([0-9A-Fa-f]{2})/g, (m, hex) =>
        String.fromCharCode(parseInt(hex, 16))
      )
    )

const quotedInput =
  deHexedInput
    .map((ol, li, f) =>
      // Escaped quote (\")
      ol.replace(/\\"/g, '"')
    )

const parsedInput = quotedInput.map((nl, li, f) => ({
  i: li,
  ol: initialInput[li],
  olLen: initialInput[li].length,
  nl,
  nlLen: nl.length,
  diff: initialInput[li].length - nl.length
}))

const answer =
  parsedInput
    .reduce(
      (td, o) => td + o.diff,
      0
    )

console.log(answer)
