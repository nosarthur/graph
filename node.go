package graph

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type (
	// Take control of the node ID management
	nodeID uuid.UUID

	// Node represents node in the graph.
	Node interface {
		ID() nodeID
		Label() string
		Nbs() []Node
		AddNb(Node)
		RemoveNb(Node)
		String() string
	}
	nodeImpl struct {
		id    nodeID // unique
		label string
	}
)

// NewNode returns a new node with proper initialization.
// A Node is a pointer to the implementation.
func NewNode() (Node, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate node ID: %s", err)
	}
	fmt.Printf("UUIDv4: %s\n", u)
	n := nodeImpl{id: nodeID(u)}
	return &n, nil
}

func (n *nodeImpl) ID() nodeID {
	return n.id
}
func (n *nodeImpl) Label() string {
	return n.label
}
func (n *nodeImpl) AddNb(Node) {
}

func (n *nodeImpl) RemoveNb(Node) {
}

func (n *nodeImpl) String() string {
	return fmt.Sprintf("Node <%s> with label <%s>", n.id, n.label)
}

func (n *nodeImpl) Nbs() []Node {
	return nil
}
