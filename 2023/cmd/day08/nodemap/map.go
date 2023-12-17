package nodemap

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TheAlgorithms/Go/math/lcm"
)

type Map struct {
	nodes map[string]Node
}


/*
	`nodeList` should be in the following format
	```
	AAA = (BBB, CCC)
 	BBB = (DDD, EEE)
 	CCC = (ZZZ, GGG)
 	DDD = (DDD, DDD)
	EEE = (EEE, EEE)
	GGG = (GGG, GGG)
	ZZZ = (ZZZ, ZZZ)
	```
*/

func New(nodeList []string) *Map {
	nodeMap := make(map[string]Node, len(nodeList))
	for _, node := range nodeList {
		node = strings.ReplaceAll(node, " ", "")
		if len(node) > 0 {
			nodeParts := strings.Split(node, "=")
			name := nodeParts[0]
			leftRight := strings.Split(strings.TrimRight(strings.TrimLeft(nodeParts[1], "("), ")"), ",")
			nodeMap[name] = *NewNode(name, leftRight[0], leftRight[1])
		}
	}
	return &Map{nodes: nodeMap}
}

func (m *Map) FindNextNode(nodeName, direction string) (nextNodeName string, err error) {
	var node Node
	var ok bool
	if node, ok = m.nodes[nodeName]; !ok {
		return "", fmt.Errorf("node \"%v\" must exist in the node map (nodes: %v)", nodeName, m.nodes)
	}
	if nextNodeName, err = node.Move(direction); err != nil {
		return "", err
	}
	return
}

func (m *Map) FindMatchingNodeNames(namePattern string) ([]string, error) {
	nodeNames := make([]string, 0, len(m.nodes)/2)
	nameRegexp, err := regexp.Compile(namePattern)
	if err != nil {
		return nodeNames, err
	}
	for name := range m.nodes {
		if nameRegexp.MatchString(name) {
			nodeNames = append(nodeNames, name)
		}
	}
	return nodeNames, nil
}

func (m *Map) FindMultiplePathLengths(directions, startingNamePattern, endingNamePattern string) (moves int, err error) {
	startingNodeNames, err := m.FindMatchingNodeNames(startingNamePattern)
	// fmt.Println("startingNodeNames:", startingNodeNames)
	if err != nil {
		return 0, err
	}
	pathLengths := make(map[string]int, len(startingNodeNames))
	for _, nodeName := range startingNodeNames {
		pathLength, err := m.PathLength(directions, nodeName, endingNamePattern)
		if err != nil {
			return 0, err
		}
		pathLengths[nodeName] = pathLength
	}
	// fmt.Println("pathLengths:", pathLengths)
	commonEndPoint := int64(1)
	for _, length := range pathLengths {
		commonEndPoint = lcm.Lcm(commonEndPoint, int64(length))
	}
	return int(commonEndPoint), nil
}

func (m *Map) PathLength(directions, startingNodeName, endingNamePattern string) (moves int, err error) {
	turns := strings.Split(directions, "")
	t := 0 // next turn number

	nextName := startingNodeName
	endingNameRegexp := regexp.MustCompile(endingNamePattern)

	for {
		if nextName, err = m.FindNextNode(nextName, turns[t]); err != nil {
			return 0, err
		}
		if moves++; endingNameRegexp.MatchString(nextName) {
			return
		}
		if t++; t == len(turns) {
			t = 0
		}
	}
}

// func (m *Map) CountMoves(directions string, startingNodeName, endingNodeName string) (moves int, err error) {
// 	turns := strings.Split(directions, "")
// 	t := 0 // next turn number

// 	var startingNodeNamePattern *regexp.Regexp
// 	if startingNodeNamePattern, err = regexp.Compile(startingNodeName); err != nil {
// 		return 0, err
// 	}

// 	var endingNodeNamePattern *regexp.Regexp
// 	if endingNodeNamePattern, err = regexp.Compile(endingNodeName); err != nil {
// 		return 0, err
// 	}

// 	var currentNodeNameList []string
// 	if currentNodeNameList, err = m.FindMatchingNodeNames(startingNodeNamePattern); err != nil {
// 		return 0, err
// 	}

// 	iterations := 0

// 	for {
// 		var nextNodeNameList []string = make([]string, 0, len(currentNodeNameList))
// 		var endingNodeNameList []string = make([]string, 0, len(currentNodeNameList))

// 		fmt.Printf("currentNodeNameList: %v\n", currentNodeNameList)

// 		for _, currentNodeName := range currentNodeNameList {
// 			var nextNodeName string
// 			if nextNodeName, err = m.FindNextNode(currentNodeName, turns[t]); err != nil {
// 				return 0, err
// 			} else {
// 				nextNodeNameList = append(nextNodeNameList, nextNodeName)
// 				if endingNodeNamePattern.MatchString(nextNodeName) {
// 					endingNodeNameList = append(endingNodeNameList, nextNodeName)
// 				}
// 			}
// 		}

		
// 		if moves++; len(endingNodeNameList) == len(nextNodeNameList) {
// 			fmt.Printf("moves, endingNodeNameList: %v, %v\n", moves, endingNodeNameList)
// 			return moves, nil
// 		}
// 		currentNodeNameList = nextNodeNameList
// 		if t++; t == len(turns) {
// 			t = 0
// 		}
// 		if iterations ++; iterations > 1_000_000 {
// 			return 0, fmt.Errorf("after %v moves, no destination was reached\n", moves)
// 		}				
// 	}
// }
