const { readFileSync } = require('fs');

console.log(readFileSync('./input.txt', 'utf-8')
  .split(/\r?\n\r?\n/)
  .map((inventoryAsString) =>
    inventoryAsString
      .split('\n')
      .reduce(
        (totalCalories, caloriesAsString) =>
          (totalCalories += parseInt(caloriesAsString)),
        0
      )
  )
  .sort((e1, e2) => e2 - e1)
  [0]);

// console.log(`There are ${caloriesPerElf.length} elves.`);
// console.log(`The highest calorie count for one elf is ${caloriesPerElf[0]}.`);
