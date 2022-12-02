const { readFileSync } = require('fs');

const caloriesPerElf = 
    readFileSync('./input.txt', 'utf-8').
    split(/\r?\n\r?\n/).
    map(inventoryAsString =>
        inventoryAsString.
            split("\n").
            reduce((totalCalories, caloriesAsString) =>
                totalCalories += parseInt(caloriesAsString)
            , 0)
    );

const elfWithMostCalories =
    caloriesPerElf.reduce((e, calories, elf) => calories > e.calories ? { elf, calories } : e, 
    { elf: 0, calories: 0});

console.log(`Out of ${caloriesPerElf.length} elves, elf #${elfWithMostCalories.elf} has the most calories (${elfWithMostCalories.calories})`);
