package main

import (
	"fmt"
)

func main() {
	arr := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := []byte{'C', 'C', 'E', 'D'}
	lenWord := len(word)
	for i := 0; i < lenWord; i++ {
		lenRowArr := len(arr)
		lenColArr := len(arr[0])
		for j := 0; j < lenRowArr; j++ {
			for h := 0; h < lenColArr; h++ {
				if word[i] == arr[j][h] {

				}
			}
		}
	}
	fmt.Println(len(arr[0]))
	fmt.Println(arr)
}
