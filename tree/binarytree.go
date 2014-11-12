package main

import (
	"fmt"
	//"go/printer"
	"math/rand"
	"time"
)

type BTree struct {
	root *Node
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func (b *BTree) Insert(value int) (tmpNode *Node) {
	tmpNode = &Node{nil, nil, value}
	if b.root == nil {
		b.root = tmpNode
		return
	}
	b.addNode(b.root, tmpNode)
	return
}

func (b *BTree) addNode(parent *Node, node *Node) *Node {

	if parent == nil {
		return node
	}
	fmt.Println(parent.value)
	if node.value <= parent.value {
		fmt.Println("left")
		parent.left = b.addNode(parent.left, node)
	} else {
		fmt.Println("right")
		parent.right = b.addNode(parent.right, node)
	}
	return parent
}

func main() {
	b := new(BTree)
	rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 10; i++ {
	//	tmpValue := rand.Intn(10)
	//	fmt.Printf("%d,\n", tmpValue)
	//	b.Insert(tmpValue)
	//}
	b.Insert(2)
	b.Insert(1)
	//printer.Fprint(b)
	fmt.Printf("%#v", b.root.left.value)
}
