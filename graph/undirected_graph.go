package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

type Graph struct {
	vertex []byte
	edge   [][]int

	//edge   [][]bool
	//edgeNum   int
	//vertexNum int
}

func (graph *Graph) PrintGraphMatrix() {
	for i := 0; i < len(graph.edge); i++ {
		for j := 0; j < len(graph.edge); j++ {
			fmt.Printf("%d ", graph.edge[i][j])
		}
		fmt.Println()
	}
}
func CreateGraph() *Graph {
	graph := new(Graph)
	graph.vertex = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	edge := [][]byte{{'A', 'C'}, {'A', 'D'}, {'C', 'D'}, {'E', 'F'}, {'D', 'G'}}
	fmt.Println(edge)

	//建立顶点与数组下标的map数组
	mapVetex := make(map[byte]int)
	for i, v := range graph.vertex {
		mapVetex[v] = i
	}
	fmt.Println("-----------------")
	fmt.Println(mapVetex)
	//初始化graph.edge
	//graph.edge = make([][]bool, len(graph.vertex))
	graph.edge = make([][]int, len(graph.vertex))
	for i := range graph.edge {
		//graph.edge[i] = make([]bool, len(graph.vertex))
		graph.edge[i] = make([]int, len(graph.vertex))
	}

	fmt.Println(graph.edge)
	for _, item := range edge {
		v0 := item[0]
		v1 := item[1]
		fmt.Println(item)
		i := mapVetex[v0]
		j := mapVetex[v1]
		//fmt.Printf("%d,%d", i, j)
		//graph.edge[i][j] = true
		graph.edge[i][j] = 1
	}
	//graph.edge = [][]bool{}
	//graph.vertex
	return graph
}

func (graph *Graph) DFSTraverse() {
	//初始化visited函数
	visited := make([]bool, len(graph.vertex))
}
func main() {

	//fmt.Println(graph)
	graph := CreateGraph()
	graph.PrintGraphMatrix()
}
