package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNode(t *testing.T) {
	n, err := NewNode()
	require.Nil(t, err, "failed to create new node")

	assert.Zero(t, n.Label,
		fmt.Sprintf("node label is not empty: %s", n.Label))
	assert.NotZero(t, n.id, "node ID is empty")
	assert.Zero(t, len(n.Nbs()), "node has neighbors")

	// adding nil as neighbor is not allowed
	assert.NotNil(t, n.AddNb(nil))

	n2, err := NewNode()
	require.Nil(t, err, "failed to create new node")

	assert.False(t, n.IsNb(n2))

	err = n.AddNb(n2)
	assert.Nil(t, err)
	assert.True(t, n.IsNb(n2))
	assert.Equal(t, []*Node{n2}, n.Nbs())
	assert.Equal(t, []*Node{n}, n2.Nbs(), "Adjacency is mutual")

	err = n.AddNb(n2)
	assert.Nil(t, err)
	assert.Equal(t, []*Node{n2}, n.Nbs())

	// no panic
	n.RemoveNb(nil)

	n.RemoveNb(n2)
	assert.Zero(t, len(n.Nbs()))
	assert.Zero(t, len(n2.Nbs()))
	assert.False(t, n.IsNb(n2))
	assert.False(t, n2.IsNb(n2))
	assert.False(t, n2.IsNb(n))
	n.RemoveNb(n2)
	assert.Zero(t, len(n.Nbs()))
	assert.Zero(t, len(n2.Nbs()))
}

func TestNode(t *testing.T) {
}
