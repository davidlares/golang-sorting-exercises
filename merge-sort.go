package main

import "fmt"

func main(){
  arr := []int {31,41,59,26,53,58,97,93,23,84,1,100,5,23}
  fmt.Println(merge_sort(arr))
}

func merge_sort(arr []int) []int {
  // check N elements
  if len(arr) == 1 {
    // only one elements
    return arr
  }
  // spliting
  size := len(arr)
  half := size/2
  left := arr[0:half]
  right :=arr[half:size]
  // ordering
  left = merge_sort(left)
  right = merge_sort(right)

  fmt.Println("==========")
  fmt.Println(left)
  fmt.Println(right)
  fmt.Println("==========")
  // combine
  return merge(left, right)
}

func merge(left, right []int) []int {
  // check if first left > first right
  last_left := left[len(left) - 1]
  first_right := right[0]
  if(last_left <= first_right){
    return append(left, right...)
  }
  // new array
  result := make([]int,0) // slice
  var x int
  for len(left) > 0 && len(right) > 0 {
    if(left[0] <= right[0]){
      x = left[0]
      left = left[1:]
    }  else {
      x = right[0]
      right = right[1:]
    }
    result = append(result,x)
  }
  // one of two left without elements -> add them to the new
  if len(left) > 0 { result = append(result,left...) }
  if len(right) > 0 { result = append(result,right...) }

  return result
}
