package main

import (
	"fmt"
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
	edge := [][]byte{{'A', 'C'}, {'A', 'D'}, {'C', 'D'}, {'E', 'F'}, {'D', 'G'}, {'D', 'E'}}
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
		graph.edge[i][j] = 1
	}
	return graph
}

func (graph *Graph) Neighbor(i int) func() (int, bool) {
	index := 0
	return func() (pos int, ok bool) {
		for {
			if index >= len(graph.vertex) {
				return
			}

			if graph.edge[i][index] == 1 {
				pos, ok = index, true
				index++
				return
			}
			index++
			//pos, ok =
		}

	}
}

func (graph *Graph) DFSTraverse() {
	//初始化visited函数
	visited := make([]bool, len(graph.vertex))

	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			graph.DFS(i, visited)
		}
	}
}

func (graph *Graph) DFS(i int, visited []bool) {
	if i >= len(visited) || visited[i] {
		return
	}
	fmt.Printf("%c ", graph.vertex[i])
	visited[i] = true //标识为已访问

	neighbor := graph.Neighbor(i)
	for {
		n, ok := neighbor()
		if !ok {
			break
		}
		graph.DFS(n, visited)
	}
}
func main() {
	graph := CreateGraph()
	graph.PrintGraphMatrix()
	graph.DFSTraverse()
}
