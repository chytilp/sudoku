package tree

import (
	"testing"
)

func TestNodeObjectCreatedWithID(t *testing.T) {
	ID := "a1"
	n := CreateNode(ID)
	if n.ID != ID {
		t.Errorf("CreateNode creates node, ID=%s, but expected was ID=%s", n.ID, ID)
	}
	stringRepr := "<Node id=" + ID + ", data=>"
	if n.String() != stringRepr {
		t.Errorf("Created node representation = %s, but expected was = %s", n.String(),
			stringRepr)
	}
}

func TestNodeObjectCreatedParams(t *testing.T) {
	ID := "a1"
	data := "data"
	n := CreateNodeFull(NodeParams{ID: ID, Data: data})
	if n.ID != ID {
		t.Errorf("CreateNodeFull creates node, ID=%s, but expected was ID=%s", n.ID, ID)
	}
	if n.Data != data {
		t.Errorf("CreateNodeFull creates node, data=%s, but expected was data=%s", n.Data, data)
	}
	stringRepr := "<Node id=" + ID + ", data=" + data + ">"
	if n.String() != stringRepr {
		t.Errorf("Created node representation = %s, but expected was = %s", n.String(),
			stringRepr)
	}
}

func TestNodeObjectWithChild(t *testing.T) {
	ID := "a1"
	n := CreateNode(ID)
	childNode := CreateNode("a1-1")
	n.AddChild(childNode)
	if len(n.Children) != 1 {
		t.Errorf("Node %s should have one child %s, but has %d children", n, childNode,
			len(n.Children))
	}
	if !childNode.IsEqual(n.Children[0]) {
		t.Errorf("Nodes %v and %v should be equal.", childNode, n.Children[0])
	}
	if !n.IsEqual(childNode.Parent) {
		t.Errorf("Nodes %v and %v should be equal.", n, childNode.Parent)
	}
}

func TestNodeObjectWithChildren(t *testing.T) {
	n := CreateNode("a1")
	childNode1 := CreateNode("a1-1")
	childNode2 := CreateNode("a1-2")
	children := []*Node{childNode1, childNode2}
	n.AddChildren(children)
	if len(n.Children) != 2 {
		t.Errorf("Node %s should have two child , but has %d children", n, len(n.Children))
	}
	if !childNode1.IsEqual(n.Children[0]) {
		t.Errorf("Nodes %v and %v should be equal.", childNode1, n.Children[0])
	}
	if !childNode2.IsEqual(n.Children[1]) {
		t.Errorf("Nodes %v and %v should be equal.", childNode2, n.Children[1])
	}
	if !n.IsEqual(childNode1.Parent) {
		t.Errorf("Nodes %v and %v should be equal.", n, childNode1.Parent)
	}
	if !n.IsEqual(childNode2.Parent) {
		t.Errorf("Nodes %v and %v should be equal.", n, childNode2.Parent)
	}
}

func TestNodeObjectSetData(t *testing.T) {
	n := CreateNode("a1")
	expected := "<Node id=a1, data=>"
	if n.String() != expected {
		t.Errorf("Node String() method should return %s, but return %s.", expected, n.String())
	}
	n.Data = "new value"
	expected = "<Node id=a1, data=new value>"
	if n.String() != expected {
		t.Errorf("Node String() method should return %s, but return %s.", expected, n.String())
	}
}
