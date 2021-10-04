package tree

import (
	"fmt"
	"strings"
)

//Tree represents tree of nodes - composition (Node struct).
type Tree struct {
	nodes []*Node
}

//NewTree is Tree obj constructor.
func NewTree(cap int) *Tree {
	var nodes []*Node
	if cap > 0 {
		nodes = make([]*Node, 0, cap)
	} else {
		nodes = make([]*Node, 0)
	}

	treeObj := Tree{nodes: nodes}
	return &treeObj
}

//AddNode method add node to tree.
func (t *Tree) AddNode(node *Node, parentID string) error {
	duplicatedNode := t.FindNode(node.ID)
	if duplicatedNode != nil {
		return fmt.Errorf("node ID=%s already exists in tree", node.ID)
	}
	if parentID == "" {
		t.nodes = append(t.nodes, node)
	}
	parent := t.FindNode(parentID)
	if parent == nil {
		return fmt.Errorf("parent node %s was not found", parentID)
	}
	parent.AddChild(node)
	return nil
}

//FindNode method try to find node in tree by node ID.
func (t *Tree) FindNode(ID string) *Node {
	var wantedNode *Node
	for _, n := range t.nodes {
		wantedNode = t.findNodeByID(n, ID)
		if wantedNode != nil {
			return wantedNode
		}
	}
	return nil
}

//findNodeByID method try to find node by wantedNodeID in node children.
func (t *Tree) findNodeByID(node *Node, wantedNodeID string) *Node {
	for _, child := range node.Children {
		if child.ID == strings.ToLower(wantedNodeID) {
			return child
		}
		subChild := t.findNodeByID(child, wantedNodeID)
		if subChild != nil {
			return subChild
		}
	}
	return nil
}

//Siblings method returns all siblings of specified node (ID).
func (t *Tree) Siblings(ID string) []*Node {
	node := t.FindNode(ID)
	if node == nil {
		return nil
	}
	if node.Parent == nil {
		return nil
	}
	if len(node.Parent.Children) == 1 {
		return nil
	}
	siblingsCount := len(node.Parent.Children) - 1
	siblings := make([]*Node, siblingsCount, siblingsCount)
	for _, n := range node.Parent.Children {
		if !node.IsEqual(n) {
			siblings = append(siblings, n)
		}
	}
	return siblings
}

//Parent method returns parent node of specified node (ID).
func (t *Tree) Parent(ID string) *Node {
	node := t.FindNode(ID)
	if node == nil {
		return nil
	}
	return node.Parent
}

//RootNodes metods returns all root nodes of tree.
func (t *Tree) RootNodes() []*Node {
	return t.nodes
}
