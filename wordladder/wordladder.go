package main

import (
	"fmt"
)

func main() {
	dict := []string{"hot", "dot", "dog", "lot", "log"}
	start := "hit"
	end := "cog"
	dict = append(dict, start, end)
	//fmt.Println(start, end, dict)
	mapAdjacency := createAdjacencyGraph(dict, start, end)
	findLadders(mapAdjacency, start, end)
}

func findLadders(mapAdjacency map[string][]string, start, end string) {
	//visited := make([]bool, )
	visited := make(map[string]int, len(mapAdjacency))
	visited[start] = 1
	queue := make([]string, len(mapAdjacency))
	queue = append(queue, start)
	mapParent := make(map[string][]string, len(mapAdjacency))
	//mapParent[start] = nil
	bfs(mapAdjacency, visited, queue, mapParent, start, end)
	printShortestPath(mapParent, start, end)
	//fmt.Println(mapParent)
}

func printShortestPath(mapParent map[string][]string, start, end string) {
	fmt.Println(mapParent)
	stack := make([]string)
	stack = append(stack, end)
	tmp := end
	if tmp == start {
		printStack()
		return
	}
	//for mapParent[tmp] != "" {
	//	fmt.Printf("%s ", tmp)
	//	tmp = mapParent[tmp]
	//}
	//fmt.Printf("%s ", tmp)
}

func bfs(mapAdjacency map[string][]string, visited map[string]int, queue []string, mapParent map[string][]string, start string, end string) {
	var cur string
	for len(queue) > 0 {
		//出队列
		cur, queue = queue[0], queue[1:]
		//if cur == end {
		//	break
		//}
		for _, value := range mapAdjacency[cur] {
			if visited[value] == 0 {
				//入队列
				queue = append(queue, value)
				visited[value] = visited[cur] + 1
				mapParent[value] = append(mapParent[value], cur)
			} else if visited[value] > visited[cur] {
				mapParent[value] = append(mapParent[value], cur)
			}

		}
	}

}

func createAdjacencyGraph(dict []string, start, end string) (mapAdjacency map[string][]string) {
	mapAdjacency = make(map[string][]string)
	lenWord := len(start)
	for _, vi := range dict {
		for _, vj := range dict {
			count := 0
			if vi != vj {
				for k := 0; k < lenWord; k++ {
					if vi[k] == vj[k] {
						count++
					}
				}
				if count == lenWord-1 { //find adjacency
					mapAdjacency[vi] = append(mapAdjacency[vi], vj)
				}
			}

		}
	}

	//fmt.Println(mapAdjacency)
	return
}
