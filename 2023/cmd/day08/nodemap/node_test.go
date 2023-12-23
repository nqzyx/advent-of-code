package nodemap_test

import (
	"reflect"
	"testing"

	"github.com/nqzyx/advent-of-code/day08/nodemap"
)

func TestNewNode(t *testing.T) {
	type In struct{ name, left, right string }
	cases := []struct {
		in   In
		want nodemap.Node
	}{
		{In{name: "node", left: "leftNode", right: "rightNode"}, nodemap.Node{Name: "node", Left: "leftNode", Right: "rightNode"}},
		{In{name: "", left: "leftNode", right: "rightNode"}, nodemap.Node{}},
		{In{name: "node", left: "", right: "rightNode"}, nodemap.Node{}},
		{In{name: "node", left: "leftNode", right: ""}, nodemap.Node{}},
	}
	for _, c := range cases {
		in, want := c.in, c.want
		got := nodemap.NewNode(in.name, in.left, in.right)
		if !reflect.DeepEqual(got, &want) {
			t.Errorf("NewNode(%v, %v, %v) == (%v), want (%v)", in.name, in.left, in.right, got, want)
		}
	}
}

func TestMove(t *testing.T) {
	type In struct {
		turn string
		node nodemap.Node
	}
	type Want struct {
		name     string
		errIsNil bool
	}
	node := nodemap.Node{Name: "x", Left: "left", Right: "right"}
	cases := []struct {
		in   In
		want Want
	}{
		{in: In{turn: "L", node: node}, want: Want{name: "left", errIsNil: true}},
		{in: In{turn: "R", node: node}, want: Want{name: "right", errIsNil: true}},
		{in: In{turn: "X", node: node}, want: Want{name: "", errIsNil: false}},
	}
	for _, c := range cases {
		in, want := c.in, c.want
		got, err := in.node.Move(in.turn)
		if got != want.name || want.errIsNil != (err == nil) {
			t.Errorf("(%v).Move(%v) == (%v, %v), want (%v, %v)", in.node, in.turn, got, err == nil, want.name, want.errIsNil)
		}
	}
}
