package graph

import "fmt"

// Edge is an abstraction of edge
type Edge struct {
	N1 Node
	N2 Node
}

func (e Edge) String() string {
	return fmt.Sprintf("(%s, %s)", e.N1.id, e.N2.id)
}
