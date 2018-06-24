package main

import(
  "fmt"
)


func slice_chunk(sliceone []string, size int){
   var chunkslice,tempchunk []string
   var count int
  // var tempchunk []string
   for _, sliceelement := range sliceone {
     for count < size {

       tempchunk = append(tempchunk,sliceelement)
       count++
     }
     chunkslice = append(chunkslice,tempchunk...)
     tempchunk = tempchunk[:0]
   }
   fmt.Printf("Slice Chunk \n",chunkslice)
}

func main(){

     sliceone := []string{"Apple","Mango","Banana","chili","cucumber","Rose"}
     //slicetwo := []string{"Tomato","Rose","Sunflower","Apple","Mango","Banana"}

     slice_chunk(sliceone,2)
}
