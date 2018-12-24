package graph

import (
	"errors"
	"fmt"
)

// Graph implements undirected graph.
type Graph struct {
	nodes map[nodeID]*Node
	//	Edges []Edge
}

// AddNodes adds nodes to the graph.
func (g *Graph) AddNodes(nodes ...*Node) error {
	for _, n := range nodes {
		if n == nil {
			return errors.New("cannot add nil as node")
		}
		g.nodes[n.id] = n
	}
	return nil
}

// HasEdge checks the existence of an edge.
func (g *Graph) HasEdge(n1 *Node, n2 *Node) bool {
	if !g.HasNode(n1) || !g.HasNode(n2) {
		return false
	}
	return false
}

// HasNode checks the existence of a node
func (g *Graph) HasNode(n *Node) bool {
	if n == nil {
		return false
	}
	if _, ok := g.nodes[n.id]; ok {
		return true
	}
	return false
}

// AddEdge adds an edge.
func (g *Graph) AddEdge(n1 *Node, n2 *Node) error {
	//	g.Edges = append(g.Edges, Edge{n1, n2})
	if !g.HasNode(n1) {
		return fmt.Errorf("%s is not a node in the graph", n1)
	}
	if !g.HasNode(n2) {
		return fmt.Errorf("%s is not a node in the graph", n2)
	}
	if err := n1.AddNb(n2); err != nil {
		return fmt.Errorf("failed to add edge between %s and %s: %v", n1, n2, err)
	}
	return nil
}

// LabelNodes (is it needed?)
func LabelNodes(g Graph) {
}
