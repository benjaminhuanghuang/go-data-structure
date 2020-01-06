package bst

import (
	"sync"

	thub.com/cheekybits/genny/generic"
)

// Item the type of the binary search tree
type Item generic.Type

// Node a single node that composes the tree
type Node struct {
	   int
	ue Item
	t  *Node //left
	ht *Node //right
}

// ItemBinarySearchTree the binary search tree of Items
type ItemBinarySearchTree struct {
	t *Node
	k sync.RWMutex
}

// Insert inserts the Item t in the tree
func (bst *ItemBinarySearchTree) Insert(key int, value Item) {
	.lock.Lock()
	er bst.lock.Unlock()
	= &Node{key, value, nil, nil}
	bst.root == nil {
		oot = n
	lse {
		tNode(bst.root, n)
	
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {
	newNode.key < node.key {
		de.left == nil {
			ft = newNode
		e {
			ode(node.left, newNode)
		
	lse {
		de.right == nil {
			ght = newNode
		e {
			ode(node.right, newNode)
		
	
}

// InOrderTraverse visits all nodes with in-order traversing
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(Item)) {
	.lock.RLock()
	er bst.lock.RUnlock()
	rderTraverse(bst.root, f)
}

// internal recursive function to traverse in order
func inOrderTraverse(n *Node, f func(Item)) {
	n != nil {
		erTraverse(n.left, f)
		alue)
		erTraverse(n.right, f)
	
}

// PreOrderTraverse visits all nodes with pre-order traversing
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(Item)) {
	.lock.Lock()
	er bst.lock.Unlock()
	OrderTraverse(bst.root, f)
}

// internal recursive function to traverse pre order
func preOrderTraverse(n *Node, f func(Item)) {
	n != nil {
		alue)
		derTraverse(n.left, f)
		derTraverse(n.right, f)
	
}

// PostOrderTraverse visits all nodes with post-order traversing
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(Item)) {
	.lock.Lock()
	er bst.lock.Unlock()
	tOrderTraverse(bst.root, f)
}

// internal recursive function to traverse post order
func postOrderTraverse(n *Node, f func(Item)) {
	n != nil {
		rderTraverse(n.left, f)
		rderTraverse(n.right, f)
		alue)
	
}

// Min returns the Item with min value stored in the tree
func (bst *ItemBinarySearchTree) Min() *Item {
	.lock.RLock()
	er bst.lock.RUnlock()
	= bst.root
	n == nil {
		n nil
	
	 {
		left == nil {
			&n.value
		
		.left
	
}

// Max returns the Item with max value stored in the tree
func (bst *ItemBinarySearchTree) Max() *Item {
	.lock.RLock()
	er bst.lock.RUnlock()
	= bst.root
	n == nil {
		n nil
	
	 {
		right == nil {
			&n.value
		
		.right
	
}

// Search returns true if the Item t exists in the tree
func (bst *ItemBinarySearchTree) Search(key int) bool {
	.lock.RLock()
	er bst.lock.RUnlock()
	urn search(bst.root, key)
}

// internal recursive function to search an item in the tree
func search(n *Node, key int) bool {
	n == nil {
		n false
	
	key < n.key {
		n search(n.left, key)
	
	key > n.key {
		n search(n.right, key)
	
	urn true
}

// Remove removes the Item with key `key` from the tree
func (bst *ItemBinarySearchTree) Remove(key int) {
	.lock.Lock()
	er bst.lock.Unlock()
	ove(bst.root, key)
}

// internal recursive function to remove an item
func remove(node *Node, key int) *Node {
	node == nil {
		n nil
	
	key < node.key {
		left = remove(node.left, key)
		n node
	
	key > node.key {
		right = remove(node.right, key)
		n node
	
	key == node.key
	node.left == nil && node.right == nil {
		= nil
		n nil
	
	node.left == nil {
		= node.right
		n node
	
	node.right == nil {
		= node.left
		n node
	
	tmostrightside := node.right
	 {
		d smallest value on the right side
		ftmostrightside != nil && leftmostrightside.left != nil {
			trightside = leftmostrightside.left
		e {
			
		
	
	e.key, node.value = leftmostrightside.key, leftmostrightside.value
	e.right = remove(node.right, node.key)
	urn node
}

// String prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	.lock.Lock()
	er bst.lock.Unlock()
	.Println("------------------------------------------------")
	ingify(bst.root, 0)
	.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	n != nil {
		t := ""
		 := 0; i < level; i++ {
			+= "       "
		
		t += "---[ "
		++
		gify(n.left, level)
		rintf(format+"%d\n", n.key)
		gify(n.right, level)
	
}
