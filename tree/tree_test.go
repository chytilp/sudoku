package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func createTree() *Tree {
	treeObj := NewTree(3)
	treeObj.AddNode(CreateNode("a"), "")
	treeObj.AddNode(CreateNode("b"), "")
	treeObj.AddNode(CreateNode("c"), "")
	treeObj.AddNode(CreateNode("a1"), "a")
	treeObj.AddNode(CreateNode("a2"), "a")
	treeObj.AddNode(CreateNode("a11"), "a1")
	treeObj.AddNode(CreateNode("a12"), "a1")
	treeObj.AddNode(CreateNode("a111"), "a11")
	treeObj.AddNode(CreateNode("a112"), "a11")
	treeObj.AddNode(CreateNode("a113"), "a11")
	return treeObj
}

func TestTreeDisplay(t *testing.T) {
	treeObj := createTree()
	representation := treeObj.Display(" - ")
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

func TestTreeShouldSayIfNodeExists(t *testing.T) {
	treeObj := createTree()
	var tests = []struct {
		id    string
		found bool
	}{
		{"a", true},
		{"b", true},
		{"c", true},
		{"d", false},
		{"a1", true},
		{"a2", true},
		{"a3", false},
		{"a11", true},
		{"a12", true},
		{"a13", false},
		{"a111", true},
		{"a112", true},
		{"a113", true},
		{"a114", false},
	}
	for _, test := range tests {
		found := treeObj.ExistsNode(test.id)
		if found != test.found {
			t.Errorf("ExistsNode for id: %s should return %t, but return %t",
				test.id, test.found, found)
		}
	}
	expectedKeys := 10
	if len(treeObj.ids) != expectedKeys {
		t.Errorf("treeObj.ids should have %d, but have %d", expectedKeys,
			len(treeObj.ids))
	}
}

func TestTreeShouldFindNode(t *testing.T) {
	expectedID := "a113"
	treeObj := createTree()
	node := treeObj.FindNode(expectedID)
	if node == nil {
		t.Fatal("findNode should return node, but return nil")
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

func TestTreeShouldReturnSiblings(t *testing.T) {
	expected := []string{"a112", "a113"}
	treeObj := createTree()
	nodeID := "a111"
	nodes := treeObj.Siblings(nodeID)
	expectedCount := 2
	if len(nodes) != expectedCount {
		t.Errorf("Siblings of node %s should be %d, but was %v", nodeID, expectedCount,
			nodes)
	}
	ids := []string{nodes[0].ID, nodes[1].ID}
	if diff := cmp.Diff(ids, expected); diff != "" {
		t.Errorf("Siblings returns: %s\n, but expected: %s\n, diff: %s\n", ids, expected, diff)
	}
}
