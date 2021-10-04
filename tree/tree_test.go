package tree

func createTree() *Tree {
	treeObj := NewTree(3)
	treeObj.AddNode(CreateNode("a"), "")
	treeObj.AddNode(CreateNode("b"), "")
	treeObj.AddNode(CreateNode("c"), "")
	treeObj.AddNode(CreateNode("a1"), "a")
	treeObj.AddNode(CreateNode("a2"), "a")
	treeObj.AddNode(CreateNode("a11"), "a1")
	treeObj.AddNode(CreateNode("a12"), "a1")
	treeObj.AddNode(CreateNode("a111"), "a11")
	treeObj.AddNode(CreateNode("a112"), "a11")
	treeObj.AddNode(CreateNode("a113"), "a11")
	return treeObj
}
