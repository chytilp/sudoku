package tree

import (
	"fmt"
	"strings"
)

//Node represents one node in solution path.
type Node struct {
	Children []*Node
	ID       string
	Parent   *Node
	Data     string
}

//NodeParams represents constructor parameters
type NodeParams struct {
	ID   string
	Data string
}

//-- Constructors --

//CreateNodeFull creates new node with specified parameters.
func CreateNodeFull(params NodeParams) *Node {
	ID := strings.TrimSpace(params.ID)
	var node Node
	if params.Data == "" {
		node = Node{ID: strings.ToLower(ID)}
	} else {
		node = Node{ID: strings.ToLower(ID), Data: params.Data}
	}

	return &node
}

//CreateNode creates new node with specified ID.
func CreateNode(ID string) *Node {
	return CreateNodeFull(NodeParams{ID: ID})
}

// -- Public methods --
//String returns node representation.
func (n Node) String() string {
	return fmt.Sprintf("<Node id=%s, data=%s >", n.ID, n.Data)
}

//AddChild method adds new node to child collection.
func (n *Node) AddChild(childNode *Node) {
	childNode.Parent = n
	n.Children = append(n.Children, childNode)
}

//AddChildren methods adds more new nodes.
func (n *Node) AddChildren(children []*Node) {
	for _, child := range children {
		n.AddChild(child)
	}
}
