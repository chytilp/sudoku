package engine

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
	Data     string
	Parent   *Node
}

//Plan represents all found solutions of 1 game.
type Plan struct {
	Paths []Node
}

//FindSolutions method searches poseible solution paths for game.
func (p *Plan) FindSolutions() (*bool, error) {
	return nil, nil
}
