package tree

import "fmt"

type Node struct {
	value any
	left  *Node
	right *Node
}

func preorder(t *Node) {
	if t == nil {
		return
	}

	fmt.Printf(" (%d) ", t.value)
	preorder(t.left)
	preorder(t.right)
}

func inorder(t *Node) {
	if t == nil {
		return
	}

	inorder(t.left)
	fmt.Printf(" (%d) ", t.value)
	inorder(t.right)
}

func postorder(t *Node) {
	if t == nil {
		return
	}

	postorder(t.left)
	postorder(t.right)
	fmt.Printf(" (%d) ", t.value)
}

func levelorder(t *Node) {
	queue := []*Node{t}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		fmt.Printf(" (%d) ", curr.value)

		if curr.left != nil {
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
}
