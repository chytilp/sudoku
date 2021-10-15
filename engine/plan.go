package engine

import (
	"github.com/chytilp/sudoku/tree"
)

const (
	doneTrue  = "done: true"
	doneFalse = "done: false"
)

//Plan represents all found solutions of 1 game.
type Plan struct {
	solutionTree *tree.Tree
	current      *tree.Node
	orderNum     byte
}

//NewPlan create instance of Plan object.
func NewPlan() *Plan {
	return &Plan{
		solutionTree: tree.NewTree(10),
	}
}

//AddNodes method add nodes to plan.
func (p *Plan) AddNodes(solutionNode string, otherNodes []string) {
	n := p.createNode(solutionNode, true)
	parentID := ""
	if p.current != nil {
		parentID = p.current.ID
	}
	p.solutionTree.AddNode(n, parentID)
	p.current = n
	for _, nID := range otherNodes {
		otherNode := p.createNode(nID, false)
		p.solutionTree.AddNode(otherNode, parentID)
	}
}

//FindNearest method returns nearest undone node.
func (p *Plan) FindNearest() *tree.Node {
	if p.current == nil {
		return nil
	}
	if !p.isNodeDone(p.current) {
		return p.current
	}
	var parent *tree.Node
	if p.current.Parent != nil {
		parent = p.current.Parent
	}
	return p.findNearestRecursive(parent)
}

//FindNode method finds node by its ID.
func (p *Plan) FindNode(ID string) *tree.Node {
	return p.solutionTree.FindNode(ID)
}

//SetCurrent method set node as current and set it as done:true
func (p *Plan) SetCurrent(node *tree.Node) {
	node.Data = doneTrue
	p.current = node
}

func (p *Plan) findNearestRecursive(node *tree.Node) *tree.Node {
	var children []*tree.Node
	if node == nil {
		children = p.solutionTree.RootNodes()
	} else {
		children = node.Children
	}
	for _, child := range children {
		if !p.isNodeDone(child) {
			return child
		}
	}
	if node == nil {
		return nil
	}
	return p.findNearestRecursive(node.Parent)
}

func (p *Plan) isNodeDone(node *tree.Node) bool {
	return node.Data == doneTrue
}

func (p *Plan) createNode(id string, done bool) *tree.Node {
	data := doneFalse
	if done {
		data = doneTrue
	}
	return tree.CreateNodeFull(tree.NodeParams{ID: id, Data: data})
}

//SetNodesDone method set data to done in certain nodes.
func (p *Plan) SetNodesDone(doneNodes []string) {
	p.setNodesUndoneRecursive(nil)
	for _, nID := range doneNodes {
		n := p.FindNode(nID)
		n.Data = doneTrue
	}
}

func (p *Plan) setNodesUndoneRecursive(node *tree.Node) {
	var children []*tree.Node
	if node == nil {
		children = p.solutionTree.RootNodes()
	} else {
		children = node.Children
	}
	if children == nil {
		return
	}
	for _, child := range children {
		child.Data = doneFalse
		p.setNodesUndoneRecursive(child)
	}
}

/*
//FindSolutions method searches for possible solution paths for game.
func (p *Plan) FindSolutions(game *structures.Game) (*bool, error) {
	//original := *game
	engine := Engine{game: game}
	err := p.findOneLevelSolutions(engine, "")
	if err != nil {
		return nil, err
	}
	ok := true
	return &ok, nil
}

func (p *Plan) findOneLevelSolutions(engine Engine, parentNode string) error {
	bestCandidates, err := engine.SelectBestCandidates()
	if err != nil {
		return err
	}
	for cellID, values := range bestCandidates {
		for _, value := range values {
			c := tree.CreateNode(cellID + "=" + fmt.Sprintf("%d", value))
			if err := p.solutionTree.AddNode(c, parentNode); err != nil {
				return err
			}
		}
	}
	return nil
}
*/
