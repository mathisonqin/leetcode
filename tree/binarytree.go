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

	if node.value <= parent.value {
		//fmt.Println("left")
		parent.left = b.addNode(parent.left, node)
	} else {
		//fmt.Println("right")
		parent.right = b.addNode(parent.right, node)
	}
	return parent
}

func Traverse(node *Node) {
	if node == nil {
		return
	}
	//fmt.Printf("%d\t", node.value)
	Traverse(node.left)
	Traverse(node.right)
}

func TraverseNonRec(node *Node) {

}

func main() {
	b := new(BTree)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		tmpValue := rand.Intn(10)
		fmt.Printf("%d\t", tmpValue)
		b.Insert(tmpValue)
		//b.Insert(i)
	}
	start := time.Now().UnixNano()
	fmt.Println("\n")
	fmt.Println("----------")
	//Traverse(b.root)
	TraverseNonRec(b.root)
	end := time.Now().UnixNano()
	//fmt.Println(time.Second)
	cost := float64((end - start)) / float64(time.Second.Nanoseconds())
	fmt.Println(cost)
	//printer.Fprint(b)
	//fmt.Printf("%#v", b.root.left.value)
}
