package engine

import (
	"testing"
)

func createPlan() *Plan {
	plan := Plan{}
	plan.AddNode(CreateNode("a"), "", "")
	plan.AddNode(CreateNode("b"), "", "")
	plan.AddNode(CreateNode("c"), "", "")
	plan.AddNode(CreateNode("a1"), "a", "")
	plan.AddNode(CreateNode("a2"), "a", "")
	plan.AddNode(CreateNode("a11"), "a", "a1")
	plan.AddNode(CreateNode("a12"), "a", "a1")
	plan.AddNode(CreateNode("a111"), "a", "a11")
	plan.AddNode(CreateNode("a112"), "a", "a11")
	plan.AddNode(CreateNode("a113"), "a", "a11")
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
	expected :=
		`a - a1 - a11 - a111
                - a112
                - a113
          - a12
   b
   c`
	if representation != expected {
		t.Errorf("Plan display returns: %s, but expected: %s", representation, expected)
	}
}
