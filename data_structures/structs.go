package data_structures

// Linked List Implementation

type Node struct{
	val int
	previous *Node
}

var start *Node
func Add(elem int) {
	start = &(Node{val:elem,previous:start})
}

func Search(elem int) *Node {
	temp := start
	for{
		if temp==nil || temp.val==elem{
			return temp
		}
		temp = temp.previous
	}
}
