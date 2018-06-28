package main

import (
	"fmt"
)

//find the min
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//func to chunk the slice
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
  fmt.Println("Array Chunk \n",chunkslice)
}

//find difference between two slice
func slicedifference(sliceone []string, slicetwo[]string) {
     var diff []string
     //loop two time to find differnence
     for i := 0; i < 2; i++ {
     for _, sliceoneelement := range sliceone {
       found := false
       for _, slicetwoelement := range slicetwo {
          if sliceoneelement == slicetwoelement{
            found = true
            break;
          }
       }
       if !found {
         diff = append(diff, sliceoneelement)
       }
     }
     //swap slice if it was first loop or i == 0
     if i == 0 {
       sliceone,slicetwo = slicetwo,sliceone
     }
   }
   fmt.Println("Array Difference\n",diff);
}

//combine two slice use one as key other as value
func slicecombine(sliceone []string, slicetwo []string){
   combined := make(map[string]string)
   if len(sliceone) != len(slicetwo){
     fmt.Println("Please use slice of same size")
     return
   } else {
      // comebine the slice
       for i :=0; i < len(sliceone); i++ {
         combined[sliceone[i]] = slicetwo[i]
       }
       fmt.Println("Combined Slice\n",combined)
   }
   return
}

//count the occurance of element in slice
func slicecount(sliceone []string, slicetwo []string){
  resultslice := make(map[string]int)
  for i := 0; i < 2; i++ {
     for _, sliceelement := range sliceone {
       resultslice[sliceelement]++
     }
     //if first loop completed i == 0 swap the slice
     if i == 0 {
       sliceone,slicetwo = slicetwo,sliceone
     }
  }
  fmt.Println("Element Count\n",resultslice)
}

//merge two slice in one
func slicemerge(sliceone []string, slicetwo []string){
   var slicemerge []string
   for i:=0; i<2; i++{
      for _, sliceelement :=  range sliceone {
         slicemerge = append(slicemerge,sliceelement)
      }
      if i==0 {
        sliceone,slicetwo = slicetwo,sliceone
      }
   }
   fmt.Println("Slice Merged\n",slicemerge)
}

//check if element available in slice
func inslice(sliceone []string, checkstring string){
  var elementfound bool = false
  for _, sliceelement := range sliceone {
     if sliceelement == checkstring {
       elementfound = true
       break
     }
  }
  if elementfound {
    fmt.Println("Element found in slice\n")
  } else {
    fmt.Println("Element not found in slice\n")
  }
}

func main() {
	sliceone := []string{"Apple", "Mango", "Banana", "chili", "cucumber", "Rose"}
	slicetwo := []string{"Tomato","Rose","Sunflower","Apple","Mango","Banana"}
  fmt.Println("Slice One\n",sliceone)
  fmt.Println("Slice Two\n",slicetwo)

  slicechunk(sliceone, 2)
  fmt.Println("-------------------------------------")
  slicedifference(sliceone,slicetwo)
  fmt.Println("-------------------------------------")
  slicecombine(sliceone,slicetwo)
  fmt.Println("-------------------------------------")
  slicecount(sliceone,slicetwo)
  fmt.Println("-------------------------------------")
  slicemerge(sliceone,slicetwo)
  fmt.Println("-------------------------------------")
  inslice(sliceone,"Apple")
}
