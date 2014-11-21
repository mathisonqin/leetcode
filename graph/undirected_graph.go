package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

type Graph struct {
	vertex []byte
	edge   [][]bool
	//edgeNum   int
	//vertexNum int
}

func CreateGraph() *Graph {
	graph := new(Graph)
	graph.vertex = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	edge := [][]byte{{'A', 'C'}, {'A', 'D'}, {'C', 'D'}, {'E', 'F'}, {'D', 'G'}}
	fmt.Println(edge)

	//建立顶点与数组下标的map数组
	mapVetex := make(map[byte]int)

	//初始化graph.edge
	graph.edge = make([][]bool, len(graph.vertex))
	for i := range graph.edge {
		graph.edge[i] = make([]bool, 2)
	}

	fmt.Println(graph.edge)
	for _, item := range edge {
		i := item[0]
		j := item[1]
		fmt.Println(item)
		for pos, vex := range graph.vertex {
			if item == vex {
				graph.edge
			}
		}
	}
	//graph.edge = [][]bool{}
	//graph.vertex
	return graph
}
func main() {

	//fmt.Println(graph)
	fmt.Println(CreateGraph())
}
