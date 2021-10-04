package engine

import (
	"fmt"
	"testing"

	"github.com/chytilp/sudoku/structures"
	"github.com/chytilp/sudoku/tree"
	"github.com/google/go-cmp/cmp"
)

func createPlan() *Plan {
	plan := Plan{}
	plan.AddNode(tree.CreateNode("a"), "", "")
	plan.AddNode(tree.CreateNode("b"), "", "")
	plan.AddNode(tree.CreateNode("c"), "", "")
	plan.AddNode(tree.CreateNode("a1"), "a", "")
	plan.AddNode(tree.CreateNode("a2"), "a", "")
	plan.AddNode(tree.CreateNode("a11"), "a", "a1")
	plan.AddNode(tree.CreateNode("a12"), "a", "a1")
	plan.AddNode(tree.CreateNode("a111"), "a", "a11")
	plan.AddNode(tree.CreateNode("a112"), "a", "a11")
	plan.AddNode(tree.CreateNode("a113"), "a", "a11")
	return &plan
}

func TestPlanShouldFindNode(t *testing.T) {
	expectedID := "a113"
	plan := createPlan()
	node, err := plan.FindNode("a", expectedID)
	if err != nil {
		t.Errorf("FindNode should return node, but appears error: %v", err)
	}
	if node.ID != expectedID {
		t.Errorf("FindNode returns node: %s, but expected nodeID: %s", node, expectedID)
	}
	expectedParentID := "a11"
	if node.Parent.ID != expectedParentID {
		t.Errorf("FindNode parent node id: %s, but expected nodeID: %s", node.Parent.ID, expectedParentID)
	}
	expectedParentParentID := "a1"
	if node.Parent.Parent.ID != expectedParentParentID {
		t.Errorf("FindNode parent parent node id: %s, but expected nodeID: %s", node.Parent.Parent.ID,
			expectedParentParentID)
	}
	expectedParentParentParentID := "a"
	if node.Parent.Parent.Parent.ID != expectedParentParentParentID {
		t.Errorf("FindNode parent parent parent node id: %s, but expected nodeID: %s",
			node.Parent.Parent.Parent.ID, expectedParentParentParentID)
	}
}

func TestPlanShouldBeCorrectlyDisplayed(t *testing.T) {
	plan := createPlan()
	representation := plan.Display(" - ")
	expected := "a - a1 - a11 - a111\n" +
		"a - a1 - a11 - a112\n" +
		"a - a1 - a11 - a113\n" +
		"a - a1 - a12\n" +
		"a - a2\n" +
		"b\n" +
		"c"
	if diff := cmp.Diff(representation, expected); diff != "" {
		t.Errorf("Plan display returns: %s\n, but expected: %s\n, diff: %s\n", representation, expected, diff)
	}
}

func TestPlanFindSolutions(t *testing.T) {
	plan := Plan{}
	g, err := structures.NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	ok, err := plan.FindSolutions(g)
	if err != nil {
		t.Errorf("FindSolutions should not return error: %v", err)
	}
	if !*ok {
		t.Errorf("FindSolutions should return ok.")
	}
	for _, n := range plan.Paths {
		fmt.Printf("%s", n.String())
	}
}
