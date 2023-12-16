package nodemap

import (
	"fmt"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func NewNode(left, right string) *Node {
	return &Node{
		Left:  left,
		Right: right,
	}
}

type Map struct {
	nodes map[string]Node
}

// For example
// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)
func NewMap(nodeList []string) *Map {
	nodeMap := make(map[string]Node, len(nodeList))
	for _, node := range nodeList {
		node = strings.ReplaceAll(node, " ", "")
		if len(node) > 0 {
			nodeParts := strings.Split(node, "=")
			name := nodeParts[0]
			leftRight := strings.Split(strings.TrimRight(strings.TrimLeft(nodeParts[1], "("), ")"), ",")
			nodeMap[name] = *NewNode(leftRight[0], leftRight[1])
		}
	}
	return &Map{nodes: nodeMap}
}

func (m *Map) CountMoves(startNode, endNode, directions string) (moves int, err error) {
	var t int
	var node Node
	var ok bool
	nextNodeName := startNode
	turns := strings.Split(directions, "")
	for {
		if node, ok = m.nodes[nextNodeName]; !ok {
			return 0, fmt.Errorf("node, \"%v\" must exist in the map", startNode)
		}
		switch turns[t] {
		case "L":
			nextNodeName = node.Left
		case "R":
			nextNodeName = node.Right
		default:
			return 0, fmt.Errorf("direction to turn \"%v\" is invalid", turns[t])
		}
		if moves++; nextNodeName == endNode {
			return
		}
		if t++; t == len(turns) {
			t = 0
		}
	}
}
