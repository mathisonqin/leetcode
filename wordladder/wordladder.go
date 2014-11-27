package main

import (
	"fmt"
)

func main() {
	dict := []string{"hot", "dot", "dog", "lot", "log"}
	start := "hit"
	end := "cog"
	dict = append(dict, start, end)
	fmt.Println(start, end, dict)
	createAdjacencyGraph(dict, start, end)
}

func createAdjacencyGraph(dict []string, start, end string) {
	mapAdjacency := make(map[string][]string)
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

	fmt.Println(mapAdjacency)
}
