package nodemap

import (
	"fmt"
	"regexp"
	"strings"
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
		return "", fmt.Errorf("node, \"%v\" must exist in the map", nodeName)
	}
	if nextNodeName, err = node.Move(direction); err != nil {
		return "", err
	}
	return
}

// func (m *Map) CountMoves(startingNodeName, endingNodeName, directions string) (moves int, err error) {
// 	turns := strings.Split(directions, "")
// 	t := 0 // next turn number

// 	nextNodeName := startingNodeName

// 	for {
// 		if nextNodeName, err = m.FindNextNode(nextNodeName, turns[t]); err != nil {
// 			return 0, err
// 		}
// 		if moves++; nextNodeName == endingNodeName {
// 			return
// 		}
// 		if t++; t == len(turns) {
// 			t = 0
// 		}
// 	}
// }

func (m *Map) FindMatchingNodeNames(r *regexp.Regexp) (nodeNames []string, err error) {
	nodeNames = make([]string, 0, len(m.nodes)/2)
	for name := range m.nodes {
		if r.MatchString(name) {
			nodeNames = append(nodeNames, name)
		}
	}
	return
}

func (m *Map) CountMoves(directions string, startingNodeName, endingNodeName string) (moves int, err error) {
	turns := strings.Split(directions, "")
	t := 0 // next turn number

	var startingNodeNamePattern *regexp.Regexp
	if startingNodeNamePattern, err = regexp.Compile(startingNodeName); err != nil {
		return 0, err
	}

	var endingNodeNamePattern *regexp.Regexp
	if endingNodeNamePattern, err = regexp.Compile(endingNodeName); err != nil {
		return 0, err
	}

	var currentNodeNameList []string

	if currentNodeNameList, err = m.FindMatchingNodeNames(startingNodeNamePattern); err != nil {
		return 0, err
	}

	for {
		var nextNodeNameList []string = make([]string, 0, len(currentNodeNameList))
		var endingNodeNameList []string = make([]string, 0, len(currentNodeNameList))
		for _, currentNodeName := range currentNodeNameList {
			var nextNodeName string
			if nextNodeName, err = m.FindNextNode(currentNodeName, turns[t]); err != nil {
				return 0, err
			} else {
				nextNodeNameList = append(nextNodeNameList, nextNodeName)
				if endingNodeNamePattern.MatchString(nextNodeName) {
					endingNodeNameList = append(endingNodeNameList, nextNodeName)
				}
			}
		}
		if moves++; len(endingNodeNameList) == len(nextNodeNameList) {
			return
		}
		currentNodeNameList = nextNodeNameList
		if t++; t == len(turns) {
			t = 0
		}
	}
}
