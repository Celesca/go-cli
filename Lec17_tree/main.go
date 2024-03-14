package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val int) {
	newNode := &Node{val: val}
	if t.root == nil {
		t.root = newNode
		return
	}

	cur := t.root
	for {
		if val < cur.val {
			if cur.left == nil {
				cur.left = newNode
				return
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = newNode
				return
			}
			cur = cur.right
		}
	}
}

func (t *Tree) Search(val int) *Node {
	cur := t.root
	for cur != nil {
		if cur.val == val {
			return cur
		}
		if val < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

func (t *Tree) InorderTraversal() {
	t.inorderTraversal(t.root)
}

func (t *Tree) inorderTraversal(node *Node) {
	if node == nil {
		return // If it is empty root
	}

	t.inorderTraversal(node.left)
	fmt.Printf("%d ", node.val)
	t.inorderTraversal(node.right)
}

func (t *Tree) PreorderTraversal() {
	t.preorderTraversal(t.root)
}

func (t *Tree) preorderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.val)
	t.preorderTraversal(node.left)
	t.preorderTraversal(node.right)
}

func (t *Tree) PostorderTraversal() {
	t.postorderTraversal(t.root)
}

func (t *Tree) postorderTraversal(node *Node) {

	if node == nil {
		return
	}

	t.postorderTraversal(node.left)
	t.postorderTraversal(node.right)
	fmt.Printf("%d ", node.val)
}

func main() {
	t := &Tree{}
	t.Insert(10)
	t.Insert(5)
	t.Insert(15)
	t.Insert(8)
	t.Insert(3)
	t.Insert(20)
	t.Insert(7)
	t.Insert(13)

	fmt.Println("\nInorderTraversal")
	t.InorderTraversal()
	fmt.Println("\nPreorderTraversal")
	t.PreorderTraversal()
	fmt.Println("\nPostorderTraversal")
	t.PostorderTraversal()
}
