package graph

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type (
	// Take control of the node ID management
	nodeID struct {
		uuid.UUID
	}

	// Node is a node in undirected graph.
	Node struct {
		id    nodeID
		nbs   map[nodeID]*Node
		Label string
		Value interface{}
	}
)

// NewNode returns a pointer to a new node with proper initialization.
func NewNode() (*Node, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate node ID: %s", err)
	}
	n := Node{
		id:  nodeID{u},
		nbs: map[nodeID]*Node{},
	}
	fmt.Println(&n)
	return &n, nil
}

// IsNb checks whether m is a neighbor of n.
func (n *Node) IsNb(m *Node) bool {
	_, ok := n.nbs[m.id]
	return ok
}

// AddNb makes m and n neighbors.
func (n *Node) AddNb(m *Node) error {
	if m == nil {
		return errors.New("cannot add nil as neighbor")
	}
	if _, ok := n.nbs[m.id]; ok {
		return nil
	}
	n.nbs[m.id] = m
	if err := m.AddNb(n); err != nil {
		return err
	}
	return nil
}

// RemoveNb makes n and m not neighbors.
func (n *Node) RemoveNb(m *Node) {
	// FIXME: throw error?
	if m == nil {
		return
	}
	if _, ok := n.nbs[m.id]; ok {
		delete(n.nbs, m.id)
		m.RemoveNb(n)
	}
}

// Nbs returns n's neighbors.
func (n *Node) Nbs() []*Node {
	// FIXME: potential data race?
	// 		  is it better to return n.nbs?
	nbs := make([]*Node, 0, len(n.nbs))
	for _, val := range n.nbs {
		nbs = append(nbs, val)
	}
	return nbs
}
func (n *Node) String() string {
	return fmt.Sprintf("Node <%v> with label <%s>", n.id, n.Label)
}
