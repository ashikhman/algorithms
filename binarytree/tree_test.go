package binarytree

import (
	"math/rand"
	"testing"
	"time"
)

func TestTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tree := New(1231)

	t.Log(tree)
}
