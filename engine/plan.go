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

//Plan represents all found solutions of 1 game.
type Plan struct {
	Paths []*Node
}

//AddNode insert new node into plan.
func (p *Plan) AddNode(node *Node, rootNodeID string, parentPath string) error {
	if len(rootNodeID) == 0 {
		p.Paths = append(p.Paths, node)
		return nil
	}
	var rootNode *Node
	for _, rootNode = range p.Paths {
		if rootNode.ID == rootNodeID {
			break
		}
	}
	if rootNode == nil {
		return fmt.Errorf("RootNodeID: %s was not found in plan", rootNodeID)
	}
	if len(parentPath) == 0 {
		rootNode.AddChild(node)
		return nil
	}
	parentNodes := strings.Split(parentPath, "/")
	wantedNode := p.findNodeByID(rootNode, parentNodes[0])
	if wantedNode == nil {
		return fmt.Errorf("ParentNodeID: %s was not found in plan", parentNodes[0])
	}
	wantedChildNode := p.findNodeByID(wantedNode, parentNodes[1])
	if wantedChildNode == nil {
		return fmt.Errorf("ParentNodeID: %s was not found in plan", parentNodes[1])
	}
	wantedChildNode.AddChild(node)
	return nil
}

//FindNode method searches for node by ID in branch which is expressed by rootNodeId.
func (p *Plan) FindNode(rootNodeID string, wantedNodeID string) (*Node, error) {
	var root *Node
	if len(rootNodeID) == 0 {
		for _, root = range p.Paths {
			if root.ID == strings.ToLower(wantedNodeID) {
				break
			}
		}
	}
	if root == nil {
		return nil, fmt.Errorf("RootNode Id: %s was not found in plan", rootNodeID)
	}
	return p.findNodeByID(root, wantedNodeID), nil
}

//findNodeByID method searches for node by ID in rootNode branch.
func (p *Plan) findNodeByID(rootNode *Node, wantedNodeID string) *Node {
	for _, child := range rootNode.Children {
		if child.ID == strings.ToLower(wantedNodeID) {
			return child
		}
		subChild := p.findNodeByID(child, wantedNodeID)
		if subChild != nil {
			return subChild
		}
	}
	return nil
}

//FindSolutions method searches for possible solution paths for game.
func (p *Plan) FindSolutions() (*bool, error) {
	return nil, nil
}
