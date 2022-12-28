package tree

import "testing"

var tree = &Node{
	value: 0,
	left: &Node{
		value: 1,
		left: &Node{
			value: 3,
			left:  nil,
			right: nil,
		},
		right: &Node{
			value: 4,
			left:  nil,
			right: nil,
		},
	},
	right: &Node{
		value: 2,
		left: &Node{
			value: 5,
			left:  nil,
			right: nil,
		},
		right: &Node{
			value: 6,
			left:  nil,
			right: nil,
		},
	},
}

func TestTree(t *testing.T) {
	t.Run("Pre-Order", func(t *testing.T) {
		preorder(tree)
	})

	t.Run("In-Order", func(t *testing.T) {
		inorder(tree)
	})

	t.Run("Post-Order", func(t *testing.T) {
		postorder(tree)
	})

	t.Run("Level-Order", func(t *testing.T) {
		levelorder(tree)
	})
}
