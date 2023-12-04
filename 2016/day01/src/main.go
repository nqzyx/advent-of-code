package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

func printErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func getArgs() (dataPath string) {
	dataBaseFolder := "../data"
	dataPath = path.Join(dataBaseFolder, "input", "data.txt")
	args := os.Args[1:]
	for a, arg := range args {
		if arg == "-e" {
			dataPath = path.Join(dataBaseFolder, "example", fmt.Sprintf("data%v.txt", args[a+1]))
			break
		}
	}
	return
}

func getInputData() (inputData []string) {
	dataPath := getArgs()
	inputDataAsByteArray, err := os.ReadFile(dataPath)
	if err != nil {
		printErr(err)
	}
	inputDataAsString := strings.ReplaceAll(string(inputDataAsByteArray), " ", "")
	inputData = strings.Split(inputDataAsString, ",")
	return
}

type Axis string

const (
	x Axis = "x"
	y Axis = "y"
)

type TurnDirection string

const (
	Left  TurnDirection = "L"
	Right TurnDirection = "R"
)

type Orientation string

const (
	North Orientation = "north"
	South Orientation = "south"
	East  Orientation = "east"
	West  Orientation = "west"
)

type Coordinates map[Axis]int

func (c Coordinates) String() string {
	return fmt.Sprintf("[x:%v,y:%v]", c[x], c[y])
}

type Position struct {
	coordinates Coordinates
	orientation Orientation
}

type TurnResults map[TurnDirection]Orientation

type OrientationFactor struct {
	factors     Coordinates
	turnResults TurnResults
}

type OrientationFactors map[Orientation]OrientationFactor

var orientationFactors = OrientationFactors{
	North: {factors: Coordinates{x: 0, y: +1}, turnResults: TurnResults{Right: East, Left: West}},
	South: {factors: Coordinates{x: 0, y: -1}, turnResults: TurnResults{Right: West, Left: East}},
	East:  {factors: Coordinates{x: +1, y: 0}, turnResults: TurnResults{Right: South, Left: North}},
	West:  {factors: Coordinates{x: -1, y: 0}, turnResults: TurnResults{Right: North, Left: South}},
}

var startPosition = Position{
	coordinates: Coordinates{x: 0, y: 0},
	orientation: North,
}

func (p Position) navigateWithoutRevisiting(directions []string) (newPosition Position) {
	newPosition = p
	visitedPositions := make(map[string]bool, 0)
	newPosKey := newPosition.coordinates.String()
	visitedPositions[newPosKey] = true
	for _, dir := range directions {
		turnDirection := dir[0:1]
		distanceToTravel, _ := strconv.Atoi(dir[1:])
		newOrientation := orientationFactors[newPosition.orientation].turnResults[TurnDirection(turnDirection)]
		newFactors := orientationFactors[newOrientation].factors

		for i := 0; i < distanceToTravel; i++ {
			coordinates := newPosition.coordinates
			newPosition = Position{
				orientation: newOrientation,
				coordinates: Coordinates{
					x: coordinates[x] + (1 * newFactors[x]),
					y: coordinates[y] + (1 * newFactors[y]),
				},
			}
			newPosKey = newPosition.coordinates.String()
			if _, ok := visitedPositions[newPosKey]; ok {
				return
			} else {
				visitedPositions[newPosKey] = true
			}
		}
	}
	return
}
func (p Position) navigate(directions []string) (newPosition Position) {
	newPosition = p
	for _, dir := range directions {
		turnDirection := dir[0:1]
		distanceToTravel, _ := strconv.Atoi(dir[1:])
		coordinates := newPosition.coordinates
		newOrientation := orientationFactors[newPosition.orientation].turnResults[TurnDirection(turnDirection)]
		newFactors := orientationFactors[newOrientation].factors
		newPosition = Position{
			orientation: newOrientation,
			coordinates: Coordinates{
				x: coordinates[x] + (distanceToTravel * newFactors[x]),
				y: coordinates[y] + (distanceToTravel * newFactors[y]),
			},
		}
	}
	return
}

func (origin Position) getDistanceTo(destination Position) (distance int) {
	distance = int(math.Abs(float64(origin.coordinates[x]))) + int(math.Abs(float64(destination.coordinates[x]))) +
		int(math.Abs(float64(origin.coordinates[y]))) + int(math.Abs(float64(destination.coordinates[y])))
	return
}

func partOne() (answer int) {
	directions := getInputData()
	newPosition := startPosition.navigate(directions)
	answer = newPosition.getDistanceTo(startPosition)
	return
}

func partTwo() (answer int) {
	directions := getInputData()
	newPosition := startPosition.navigateWithoutRevisiting(directions)
	answer = newPosition.getDistanceTo(startPosition)
	return
}

func main() {
	fmt.Printf("Part 1 Answer: %v\n", partOne())
	fmt.Printf("Part 2 Answer: %v\n", partTwo())
}
