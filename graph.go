package graph

import (
	"fmt"
)

// Graph is an abstraction of graph
type Graph interface {
	Nodes() []Node
	Edges() []Edge
	AddNodes(...Node) error
	AddEdge(Node, Node) error
}

type graph struct {
	nodes map[Node][]Node
	Edges []Edge
}

// AddNodes adds nodes
func (g *graph) AddNodes(nodes ...Node) {
	for _, n := range nodes {
		g.nodes[n] = []Node{}
	}
}

// AddEdge adds a new edge
func (g *graph) AddEdge(n1 Node, n2 Node) {
	g.Edges = append(g.Edges, Edge{n1, n2})
}

// LabelNodes (is it needed?)
func LabelNodes(g Graph) {
	i := 1
	for _, n := range g.Nodes() {
		//		n.label = i
		fmt.Print(n.ID())
		i++
	}
}
