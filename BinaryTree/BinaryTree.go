package BinaryTree

type BinaryTree struct {
	root *BinaryNode
}

func NewEmptyBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(key int, data string) error {
	if t.root == nil {
		t.root = NewBinaryNode(key, data)
		return nil
	}

	return t.root.Insert(key, data)
}

func (t *BinaryTree) Find(key int) (string, error) {
	return t.root.Find(key)
}

func (t *BinaryTree) Delete(key int) error {
	_, err := t.root.Delete(key)
	return err
}

func (t *BinaryTree) InOrder() {
	t.root.InOrder()
}
