package nodemap

import (
	"fmt"
)

type Node struct {
	Name  string
	Left  string
	Right string
}


func NewNode(name, left, right string) *Node {
	if name == "" || left == "" || right == "" {
		return &Node{}
	}
	return &Node{
		Name:  name,
		Left:  left,
		Right: right,
	}
}

func (n *Node) Move(direction string) (name string, err error) {
	switch direction {
	case "L":
		return n.Left, nil
	case "R":
		return n.Right, nil
	default:
		return "", fmt.Errorf("turn direction \"%v\" is invalid", direction)
	}
}
