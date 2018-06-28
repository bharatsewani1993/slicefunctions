package main

import (
	"fmt"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func slicechunk(sliceone []string, size int) {
	chunkslice := make(map[int][]string)
	var tempchunk []string
	var count int
	// var tempchunk []string
	for i := 0; i < len(sliceone); i += size {
	  tempchunk = sliceone[i:min(i+size, len(sliceone))]
    chunkslice[count] = tempchunk
    count++
  }
  fmt.Println("Chunk \n",chunkslice)
}
func main() {

	sliceone := []string{"Apple", "Mango", "Banana", "chili", "cucumber", "Rose"}
	//slicetwo := []string{"Tomato","Rose","Sunflower","Apple","Mango","Banana"}
	slicechunk(sliceone, 2)
}
