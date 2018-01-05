package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRootNode(t *testing.T) {
	g := NewGraph()
	n := Node{nil, nil}
	g.AddNode(&n)
	require.Equal(t, g.root, &n, "First node added will become root")
}
