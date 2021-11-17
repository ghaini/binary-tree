package main

import (
	"BST/BinaryTree"
	"log"
)

func main() {
	bst := BinaryTree.NewEmptyBinaryTree()
	err := bst.Insert(50, "FORD")
	if err != nil {
		log.Println(err)
		return
	}

	err = bst.Insert(60, "BENZ")
	if err != nil {
		log.Println(err)
		return
	}

	err = bst.Insert(40, "BMW")
	if err != nil {
		log.Println(err)
		return
	}

	car, err :=bst.Find(40)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(car)
	err = bst.Delete(60)
	if err != nil {
		log.Println(err)
		return
	}

	bst.InOrder()
}
