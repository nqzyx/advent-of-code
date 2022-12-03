const { readFileSync } = require('fs');

const plays = {
    Rock: { value: 1, beats: "Scissors" }, 
    Paper: { value: 2, beats: "Rock" }, 
    Scissors: { value: 3, beats: "Paper" }
};

const choices = { A: "Rock", B: "Paper", C: "Scissors" };

const outcomes = { 
    X: { value: 0, name: "Lose"}, 
    Y: { value: 3, name: "Draw"}, 
    Z: { value: 6, name: "Win" }
};


function getMyPlay(theirChoice, outcome) {
    switch (outcomes[outcome].name) {
      case "Lose":
        return plays[plays[choices[theirChoice]].beats]
        break

      case "Draw":
        return plays[choices[theirChoice]]
        break

      case "Win":
        return plays[
          Object.keys(plays).find(
            (p) => plays[p].beats === choices[theirChoice]
          )
        ]
        break

      default:
        throw new Error(`Unknown outcome (${outcome}`)
        break
    }
}

console.log(readFileSync('./input.txt', 'utf-8')
    .split("\n")
    .reduce((score, play) => {
        const [ theirChoice, outcome ] = play.split(" ");
        return (score + getMyPlay(theirChoice, outcome).value + outcomes[outcome].value);
    }, 0)
);
