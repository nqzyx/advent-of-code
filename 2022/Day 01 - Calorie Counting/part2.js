const { execFile } = require('child_process');
const { readFileSync } = require('fs');

const topElves = 
    // Array of inventory (strings) for each elf
    readFileSync('./input.txt', 'utf-8').
    split(/\r?\n\r?\n/).
    // Convert inventory (strings) to object { elf, calories }
    map((inventoryAsString, elf) => {
        return ({   
            elf, 
            calories: inventoryAsString.
                split("\n").
                reduce((totalCalories, caloriesAsString) =>
                    totalCalories += parseInt(caloriesAsString)
                , 0)
        })
    }).
    // Sort descending by calories
    sort((elf_a, elf_b) => -(elf_a.calories - elf_b.calories)).
    // Take the top 3
    slice(0,3);

const totalCalories = topElves.reduce((total, elf) => total += elf.calories, 0);

console.log(`The top ${topElves.length} elves are carrying ${totalCalories} calories`);
;