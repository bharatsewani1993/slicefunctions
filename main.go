package main

import(
  "fmt"
)


func slice_chunk(sliceone []string, size int){
   chunkslice := make(map[int][]string)
   var tempchunk []string
   var count int = 0
  // var tempchunk []string
   for i, sliceelement := range sliceone {
     if count < size {
       tempchunk = append(tempchunk,sliceelement)
       count++
     } else {
       count = 0
     }
     fmt.Printf("Temp Chunk \n",tempchunk)
     chunkslice[i] = tempchunk
     tempchunk = tempchunk[:0]
   }
   fmt.Printf("Slice Chunk %v\n",chunkslice)
}

func main(){

     sliceone := []string{"Apple","Mango","Banana","chili","cucumber","Rose"}
     //slicetwo := []string{"Tomato","Rose","Sunflower","Apple","Mango","Banana"}

     slice_chunk(sliceone,2)
}
