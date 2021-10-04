package tree

import "errors"

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
	return errors.New("")
}

//FindNode method try to find node in tree by node ID.
func (t *Tree) FindNode(ID string) *Node {
	return nil
}

//Siblings method returns all siblings of specified node (ID).
func (t *Tree) Siblings(ID string) []*Node {
	return nil
}

//Parent method returns parent node of specified node (ID).
func (t *Tree) Parent(ID string) *Node {
	return nil
}

//RootNodes metods returns all root nodes of tree.
func (t *Tree) RootNodes() []*Node {
	return t.nodes
}
