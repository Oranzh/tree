package tree

type Node interface {
	GetID() int
	GetParentID() int
	AddChild(node Node)
}

// GenerateTree  convert list to tree according to parentID
func GenerateTree(list []Node) Node {
	dataMap := make(map[int]Node, len(list))
	var rootNode Node

	for _, v := range list {
		if rootNode == nil || rootNode.GetParentID() > v.GetParentID() {
			rootNode = v
		}
		dataMap[v.GetID()] = v
	}

	for _, v := range list {
		parentNode, ok := dataMap[v.GetParentID()]
		if ok {
			parentNode.AddChild(v)
		}
	}

	return rootNode
}
