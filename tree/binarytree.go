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

func TraversePreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d\t", node.value)
	TraversePreOrder(node.left)
	TraversePreOrder(node.right)
}

func TraverseInOrder(node *Node) {
	if node == nil {
		return
	}
	TraverseInOrder(node.left)
	fmt.Printf("%d\t", node.value)
	TraverseInOrder(node.right)
}

func TraversePostOrder(node *Node) {
	if node == nil {
		return
	}
	TraversePostOrder(node.left)
	TraversePostOrder(node.right)
	fmt.Printf("%d\t", node.value)
}

func TraverseNonRec(node *Node) {
	cur := node
	stack := []*Node{}
	var tmpNode *Node
	for cur != nil || len(stack) != 0 {
		fmt.Printf("%d\t", cur.value)
		stack = append(stack, cur)
		cur = cur.left

		for cur == nil && len(stack) != 0 {

			tmpNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
			cur = tmpNode.right
		}
	}
}

func TraversePreNonRec(node *Node) {
	cur := node
	if cur == nil {
		return
	}

	stack := []*Node{}
	stack = append(stack, cur)
	for cur != nil {
		fmt.Printf("%d\t", cur.value)
		//cur = stack[lenStack-1]
		cur = cur.left
		for cur == nil && len(stack) > 0 {
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
			cur = cur.right
		}
		stack = append(stack, cur)
	}

}

func TraverseInNonRec(node *Node) {
	cur := node
	if cur == nil {
		return
	}
	stack := []*Node{}
	stack = append(stack, cur)
	for cur != nil {
		cur = cur.left
		for cur == nil && len(stack) > 0 {
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
			fmt.Printf("%d\t", cur.value)
			cur = cur.right
		}
		stack = append(stack, cur)
	}
}

func TraversePostNonRec(node *Node) {
	cur := node
	if cur == nil {
		return
	}
	type SNode struct {
		node     *Node
		rvisited bool
	}
	stack := []*SNode{}

	stack = append(stack, &SNode{node: cur, rvisited: true})
	for cur != nil {
		cur = cur.left
		for cur == nil && len(stack) > 0 {
			topSNode := stack[len(stack)-1]
			if topSNode.rvisited == true {
				topSNode.rvisited = false
				cur = topSNode.node.right
			} else {
				fmt.Printf("%d\t", topSNode.node.value)
				stack = stack[:len(stack)-1]
			}

		}
		stack = append(stack, &SNode{node: cur, rvisited: true})
	}
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
	//start := time.Now().UnixNano()
	fmt.Println("\n----------preorder-----------")
	TraversePreOrder(b.root)
	fmt.Println("\n----------inorder-----------")
	TraverseInOrder(b.root)
	fmt.Println("\n----------postorder-----------")
	TraversePostOrder(b.root)
	//TraverseNonRec(b.root)
	fmt.Println("\n----------non rec preorder-----------")
	TraversePreNonRec(b.root)
	fmt.Println("\n----------non rec inorder-----------")
	TraverseInNonRec(b.root)
	fmt.Println("\n----------non rec postorder-----------")
	TraversePostNonRec(b.root)
	fmt.Println("")
	//end := time.Now().UnixNano()
	////fmt.Println(time.Second)
	//cost := float64((end - start)) / float64(time.Second.Nanoseconds())
	//fmt.Println(cost)
	//printer.Fprint(b)
	//fmt.Printf("%#v", b.root.left.value)
}
