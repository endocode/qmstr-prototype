package model

type Graph struct {
	root  *Node
	nodes []*Node
}

type Property struct {
	Name string
}

type Node struct {
	neighbors  []*Node
	properties []*Property
}

type Edge struct {
	start       *Node
	destination *Node
}

func NewGraph() *Graph {
	return &Graph{nil, []*Node{}}
}

func (g *Graph) AddNode(newNode *Node) {
	g.nodes = append(g.nodes, newNode)
	if g.root == nil {
		g.root = newNode
	}
}
