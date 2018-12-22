package main

import "fmt"

func main(){
  arr := []int {412,12,123,1,2324,3,155,23412,2}
  fmt.Println(arr)
  ordered := selection_sort(arr)
  fmt.Println(ordered)
}

func selection_sort(arr []int) []int {
  for i:= 0; i < len(arr); i++ {
    min, pos := arr[i],i
    original := arr[i]
    for j:= i + 1; j < len(arr); j++ {
      compared := arr[j]
      if compared < min {
        min, pos = compared, j
      }
    }
    if min != original{
      // change positions
      arr[i] = min
      arr[pos] = original
    }
  }

  return arr
}
