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
	visited := make(map[string]bool, len(mapAdjacency))
	visited[start] = true
	queue := make([]string, len(mapAdjacency))
	queue = append(queue, start)
	mapParent := make(map[string]string, len(mapAdjacency))
	mapParent[start] = ""
	bfs(mapAdjacency, visited, queue, mapParent, start, end)
	printShortestPath(mapParent, start, end)
	//fmt.Println(mapParent)
}

func printShortestPath(mapParent map[string]string, start, end string) {
	tmp := end

	for mapParent[tmp] != "" {
		fmt.Printf("%s ", tmp)
		tmp = mapParent[tmp]
	}
	fmt.Printf("%s ", tmp)
}

func bfs(mapAdjacency map[string][]string, visited map[string]bool, queue []string, mapParent map[string]string, start string, end string) {
	var cur string
	for len(queue) > 0 {
		//出队列
		cur, queue = queue[0], queue[1:]
		//if cur == end {
		//	break
		//}
		for _, value := range mapAdjacency[cur] {
			if visited[value] == false {
				//入队列
				queue = append(queue, value)
				visited[value] = true
				mapParent[value] = cur
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
