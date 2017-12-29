package interview

import (
	"strconv"
	"strings"
)

// SerializeBinaryTree serializes a binary tree to a string
func SerializeBinaryTree(tree *BTree) string {

	if tree == nil {
		return ""
	}

	q := NewQueue()
	q.Add(tree.root)

	var serialized []string
	for !q.IsEmpty() {

		node := q.Remove().(*bTreeNode)
		if node.left != nil {

			q.Add(node.left)
		}
		if node.right != nil {

			q.Add(node.right)
		}

		serialized = append(
			serialized,
			strconv.Itoa(node.data),
		)
	}

	return strings.Join(serialized, ",")
}

// DeserializeBinaryTree turns a string into a binary tree
func DeserializeBinaryTree(serialized string) (*BTree, error) {

	values := strings.Split(serialized, ",")
	tree := new(BTree)

	for _, v := range values {

		value, err := strconv.Atoi(v)
		if err != nil {

			return nil, err
		}

		if tree.root == nil {

			tree.root = &bTreeNode{data: value}
		} else {

			tree.Insert(value)
		}
	}

	return tree, nil
}
