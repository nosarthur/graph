package graph

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type (
	// Take control of the node ID management
	nodeID struct{ uuid.UUID }

	// Node is an abstraction of node in undirected graph.
	Node interface {
		ID() nodeID
		Label() string
		SetLabel(string)
		Nbs() []Node
		// AddNb is idempotent
		AddNb(Node) error
		// RemoveNb is idempotent
		RemoveNb(Node)
		String() string
		IsNb(Node) bool
	}
	nodeImpl struct {
		id    nodeID // unique
		label string
		nbs   map[nodeID]Node
	}
)

// NewNode returns a new node with proper initialization.
// A Node is a pointer to the implementation.
func NewNode() (Node, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate node ID: %s", err)
	}
	n := nodeImpl{id: nodeID{u},
		nbs: map[nodeID]Node{},
	}
	fmt.Println(&n)
	return &n, nil
}

func (n *nodeImpl) ID() nodeID {
	return n.id
}
func (n *nodeImpl) IsNb(m Node) bool {
	_, ok := n.nbs[m.ID()]
	return ok
}
func (n *nodeImpl) Label() string {
	return n.label
}
func (n *nodeImpl) SetLabel(s string) {
	n.label = s
}

func (n *nodeImpl) AddNb(m Node) error {
	if m == nil {
		return errors.New("cannot add nil as neighbor")
	}
	if _, ok := n.nbs[m.ID()]; ok {
		return nil
	}
	n.nbs[m.ID()] = m
	if err := m.AddNb(n); err != nil {
		return err
	}
	return nil
}

func (n *nodeImpl) RemoveNb(m Node) {
	// FIXME: throw error?
	if m == nil {
		return
	}
	if _, ok := n.nbs[m.ID()]; ok {
		delete(n.nbs, m.ID())
		m.RemoveNb(n)
	}
}

func (n *nodeImpl) Nbs() []Node {
	// FIXME: potential data race?
	// 		  is it better to return n.nbs?
	nbs := make([]Node, 0, len(n.nbs))
	for _, val := range n.nbs {
		nbs = append(nbs, val)
	}
	return nbs
}
func (n *nodeImpl) String() string {
	return fmt.Sprintf("Node <%v> with label <%s>", n.id, n.label)
}
