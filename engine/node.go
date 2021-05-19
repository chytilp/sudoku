package engine

import (
	"fmt"
	"strconv"
	"strings"
)

//NodeType enum type.
type NodeType int

//Values of NodeType enum.
const (
	IDNode NodeType = iota
	ValueNode
)

//Node represents one node in solution path.
type Node struct {
	NodeType NodeType
	Children []*Node
	ID       string
	Parent   *Node
}

//String returns node representation.
func (n *Node) String() string {
	return fmt.Sprintf("id=%s, type=%d", n.ID, n.NodeType)
}

//CreateNode creates new node with specified ID.
func CreateNode(ID string) *Node {
	ID = strings.TrimSpace(ID)
	nodeType := IDNode
	if len(ID) == 1 {
		_, err := strconv.Atoi(ID)
		if err == nil {
			nodeType = ValueNode
		}
	}
	node := Node{NodeType: nodeType, ID: strings.ToLower(ID)}
	return &node
}

//AddChild method adds new node to child collcetion.
func (n *Node) AddChild(childNode *Node) {
	childNode.Parent = n
	n.Children = append(n.Children, childNode)
}

func (n *Node) PrintNode() string {
	return ""
}
