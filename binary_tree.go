package interview

// NewBinaryTree creates a new Binary Tree
func NewBinaryTree(data int) *BTree {

	return &BTree{
		root: &bTreeNode{
			data: data,
		},
	}
}

// BTree is a binary tree
type BTree struct {
	root *bTreeNode
}

// Insert data into the binary tree
func (btree *BTree) Insert(data int) {

	if btree == nil {
		btree = NewBinaryTree(data)
	}

	q := NewQueue()
	q.Add(btree.root)

	for !q.IsEmpty() {

		node := q.Remove().(*bTreeNode)
		if node.left != nil {

			q.Add(node.left)
		} else {

			node.left = &bTreeNode{data: data}
			break
		}

		if node.right != nil {

			q.Add(node.right)
		} else {

			node.right = &bTreeNode{data: data}
			break
		}
	}
}

// Remove an element from the binary tree
func (btree *BTree) Remove(data int) {

}

type bTreeNode struct {
	data  int
	left  *bTreeNode
	right *bTreeNode
}
