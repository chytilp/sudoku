package engine

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chytilp/sudoku/structures"
	"github.com/chytilp/sudoku/tree"
	"github.com/chytilp/sudoku/utils"
)

//Plan represents all found solutions of 1 game.
type Plan struct {
	Paths []*tree.Node
}

//AddNode insert new node into plan.
func (p *Plan) AddNode(node *tree.Node, rootNodeID string, parentPath string) error {
	//rootNodeID is empty, add node to root.
	if len(rootNodeID) == 0 {
		p.Paths = append(p.Paths, node)
		return nil
	}
	//find rootNodeID between nodes in root.
	var rootNode *tree.Node
	for _, rootNode = range p.Paths {
		if rootNode.ID == rootNodeID {
			break
		}
	}
	if rootNode == nil {
		return fmt.Errorf("RootNodeID: %s was not found in plan", rootNodeID)
	}
	//parentPath empty add to root node.
	if len(parentPath) == 0 {
		rootNode.AddChild(node)
		return nil
	}
	//TODO: simplify, it is really strange
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
func (p *Plan) FindNode(rootNodeID string, wantedNodeID string) (*tree.Node, error) {
	var root *tree.Node
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
func (p *Plan) findNodeByID(rootNode *tree.Node, wantedNodeID string) *tree.Node {
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
func (p *Plan) FindSolutions(game *structures.Game) (*bool, error) {
	//original := *game
	engine := Engine{game: game}
	err := p.findOneLevelSolutions(engine, "", "")
	if err != nil {
		return nil, err
	}
	ok := true
	return &ok, nil
}

func (p *Plan) findOneLevelSolutions(engine Engine, rootNode string, parentNode string) error {
	bestCandidates, err := engine.SelectBestCandidates()
	if err != nil {
		return err
	}
	for cellID, values := range bestCandidates {
		for _, value := range values {
			c := tree.CreateNode(cellID + "=" + fmt.Sprintf("%d", value))
			p.AddNode(c, rootNode, parentNode)
		}
	}
	return nil
}

//Display returns text represantion of tree.
func (p *Plan) Display(sep string) string {
	var result []string
	root := tree.CreateNode("")
	root.AddChildren(p.Paths)
	p.walk(root, &result, sep)
	return strings.Join(result, "\n")
}

//Auxiliary method for Display. Sort children nodes.
func (p *Plan) sortChildren(children []*tree.Node) []*tree.Node {
	ids := make([]string, len(children))
	nodes := make(map[string]*tree.Node)
	for idx, child := range children {
		ids[idx] = child.ID
		nodes[child.ID] = child
	}
	sort.Strings(ids)
	sorted := make([]*tree.Node, len(children))
	for idx, id := range ids {
		sorted[idx] = nodes[id]

	}
	return sorted
}

//Auxiliary method for Display. Walk through nodes.
func (p *Plan) walk(n *tree.Node, result *[]string, sep string) {
	children := p.sortChildren(n.Children)
	for _, child := range children {
		if child.Children == nil {
			p.printNodeLine(child, sep, result)
		} else {
			p.walk(child, result, sep)
		}
	}
}

//Auxiliary method for Display. Puts Node to output slice.
func (p *Plan) printNodeLine(n *tree.Node, sep string, output *[]string) {
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
