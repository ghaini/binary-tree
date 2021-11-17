package BinaryTree

import "fmt"

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	key   int
	data  string
}

func NewBinaryNode(key int, data string) *BinaryNode {
	return &BinaryNode{
		key:  key,
		data: data,
	}
}

func (n *BinaryNode) Insert(key int, data string) error {
	if n == nil {
		return nil
	}

	if n.key == key {
		return KeyAlreadyExistsError
	}

	if n.key > key {
		if n.left == nil {
			n.left = NewBinaryNode(key, data)
			return nil
		} else {
			return n.left.Insert(key, data)
		}
	} else {
		if n.right == nil {
			n.right = NewBinaryNode(key, data)
			return nil
		} else {
			return n.right.Insert(key, data)
		}
	}
}

func (n *BinaryNode) Find(key int) (string, error) {
	if n == nil {
		return "", KeyNotFoundError
	}

	if n.key == key {
		return n.data, nil
	}

	if n.key > key {
		return n.left.Find(key)
	} else {
		return n.right.Find(key)
	}
}

func (n *BinaryNode) Height() int {
	if n == nil {
		return 0
	}

	x := n.left.Height()
	y := n.right.Height()

	if x > y {
		return x + 1
	} else {
		return y + 1
	}

}

func (n *BinaryNode) InPredecessor() *BinaryNode {
	predecessor := n
	for predecessor.right != nil {
		predecessor = predecessor.right
	}

	return predecessor
}

func (n *BinaryNode) InSuccessor() *BinaryNode {
	successor := n
	for successor.left != nil {
		successor = successor.left
	}

	return successor
}

func (n *BinaryNode) Delete(key int) (bn *BinaryNode, err error) {
	if n == nil {
		return nil, KeyNotFoundError
	}

	if n.left == nil && n.right == nil {
		if n.key == key {
			return nil, KeyNotFoundError
		}
	}

	if n.key > key {

		n.left, err = n.left.Delete(key)
		return n, err
	}

	if n.key < key {
		n.right, err = n.right.Delete(key)
		return n, err
	}

	if n.left.Height() > n.right.Height() {
		p := n.left.InPredecessor()
		n.data = p.data
		n.key = p.key
		n.left, err = n.left.Delete(p.key)
	} else {
		s := n.right.InSuccessor()
		n.data = s.data
		n.key = s.key
		n.right, err = n.right.Delete(s.key)
	}

	return n, err
}

func (n *BinaryNode) InOrder() {
	if n == nil {
		return
	}

	n.left.InOrder()
	fmt.Print(n.key, " ")
	n.right.InOrder()
}
