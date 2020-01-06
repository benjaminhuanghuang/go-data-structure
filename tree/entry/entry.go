package main

import "../tree"

func main() {
	var root tree.Node
	root = tree.Node{value: 4}


	root.left = &tree.Node{	}
	

	fmt.Println(root)
}
