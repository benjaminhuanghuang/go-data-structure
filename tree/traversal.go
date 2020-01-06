package tree

func (node Node) Print(){
	fmt.Print(node.value, " ")
}


func (node *Node) Traverse(){
	if nood == nil
		return

	node.Left.Traverse();
	node.Print()
	node.Right.Traverse();
}