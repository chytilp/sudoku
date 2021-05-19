package engine

import (
	"fmt"
	"strings"
)

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
	parentNodePaths := strings.Split(parentPath, "/")
	previous := rootNode
	for _, parentNodePath := range parentNodePaths {
		parentNode := p.findNodeByID(previous, parentNodePath)
		if parentNode == nil {
			return fmt.Errorf("ParentNodeID: %s was not found in plan", parentNodePath)
		}
		if parentNodePath == parentNodePaths[len(parentNodePaths)-1] {
			//is it last part?
			parentNode.AddChild(node)
		} else {
			previous = parentNode
		}
	}
	return nil
}

//FindNode method searches for node by ID in branch which is expressed by rootNodeId.
func (p *Plan) FindNode(rootNodeID string, wantedNodeID string) (*Node, error) {
	var root *Node
	for _, root = range p.Paths {
		if root.ID == strings.ToLower(rootNodeID) {
			break
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

//Display returns text represantion of tree.
func (p *Plan) Display(sep string) string {
	return ""
}
