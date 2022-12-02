const { readFileSync } = require('fs');

const choices = {
    Rock: { value: 1, beats: "Scissors" }, 
    Paper: { value: 2, beats: "Rock" }, 
    Scissors: { value: 3, beats: "Paper" }
};

const theirChoices = { A: "Rock", B: "Paper", C: "Scissors" };
const myChoices = { X: "Rock", Y: "Paper", Z: "Scissors" };

function playOutcome(theirChoice, myChoice) {
    const myChoiceValue = choices[myChoices[myChoice]].value;
    let myOutcomeValue;
    if (theirChoices[theirChoice] === myChoices[myChoice]) {
        myOutcomeValue = 3
    } else if (choices[myChoices[myChoice]].beats === theirChoices[theirChoice]) {
        myOutcomeValue = 6
    } else if (choices[theirChoices[theirChoice]].beats === myChoices[myChoice]) {
        myOutcomeValue = 0
    } else {
        throw new Error(`Unknown Outcome - TheirChoice: ${theirChoice}, MyChoice: ${myChoice}`)
    }
    return myChoiceValue + myOutcomeValue;
}

const myFinalScore = readFileSync('./input.txt', 'utf-8')
    .split("\n")
    .reduce((score, play) => {
        const [ theirChoice, myChoice ] = play.split(" ");

        return (score + playOutcome(theirChoice, myChoice));
    }, 0);

console.log(`My final score: ${myFinalScore}`)