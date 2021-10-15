package tree

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chytilp/sudoku/utils"
)

//Tree represents tree of nodes - composition (Node struct).
type Tree struct {
	nodes []*Node
	ids   map[string]bool
}

//NewTree is Tree obj constructor.
func NewTree(cap int) *Tree {
	var nodes []*Node
	if cap > 0 {
		nodes = make([]*Node, 0, cap)
	} else {
		nodes = make([]*Node, 0)
	}

	treeObj := Tree{nodes: nodes, ids: make(map[string]bool)}
	return &treeObj
}

//AddNode method add node to tree.
func (t *Tree) AddNode(node *Node, parentID string) error {
	isDuplicatedNode := t.ExistsNode(node.ID)
	if isDuplicatedNode {
		return fmt.Errorf("node ID=%s already exists in tree", node.ID)
	}
	if parentID == "" {
		t.nodes = append(t.nodes, node)
		t.ids[node.ID] = true
		return nil
	}
	parent := t.FindNode(parentID)
	if parent == nil {
		return fmt.Errorf("parent node %s was not found", parentID)
	}
	parent.AddChild(node)
	t.ids[node.ID] = true
	return nil
}

//ExistsNode method returns if node (ID) exists in tree.
func (t *Tree) ExistsNode(ID string) bool {
	_, ok := t.ids[ID]
	return ok
}

//FindNode method try to find node in tree by node ID.
func (t *Tree) FindNode(ID string) *Node {
	existsNode := t.ExistsNode(ID)
	if !existsNode {
		return nil
	}
	var wantedNode *Node
	for _, n := range t.nodes {
		if n.ID == strings.ToLower(ID) {
			return n
		}
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
	idx := 0
	for _, n := range node.Parent.Children {
		if !node.IsEqual(n) {
			siblings[idx] = n
			idx++
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

//Display returns text represantion of tree.
func (t *Tree) Display(sep string) string {
	var result []string
	root := CreateNode("")
	root.AddChildren(t.nodes)
	t.walk(root, &result, sep)
	return strings.Join(result, "\n")
}

//Auxiliary method for Display. Sort children nodes.
func (t *Tree) sortChildren(children []*Node) []*Node {
	ids := make([]string, len(children))
	nodes := make(map[string]*Node)
	for idx, child := range children {
		ids[idx] = child.ID
		nodes[child.ID] = child
	}
	sort.Strings(ids)
	sorted := make([]*Node, len(children))
	for idx, id := range ids {
		sorted[idx] = nodes[id]

	}
	return sorted
}

//Auxiliary method for Display. Walk through nodes.
func (t *Tree) walk(n *Node, result *[]string, sep string) {
	children := t.sortChildren(n.Children)
	for _, child := range children {
		if child.Children == nil {
			t.printNodeLine(child, sep, result)
		} else {
			t.walk(child, result, sep)
		}
	}
}

//Auxiliary method for Display. Puts Node to output slice.
func (t *Tree) printNodeLine(n *Node, sep string, output *[]string) {
	result := []string{}
	node := n
	for node != nil && len(node.ID) > 0 {
		result = append(result, node.ID)
		node = node.Parent
	}
	utils.ReverseSlice(result)
	line := strings.Join(result, sep)
	*output = append(*output, line)
}
