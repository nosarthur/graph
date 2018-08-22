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
	assert.Zero(t, n.Label(),
		fmt.Sprintf("node label is not empty: %s", n.Label()))
	assert.NotZero(t, n.ID(), "node ID is empty")
}
