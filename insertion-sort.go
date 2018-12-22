package main

import "fmt"

func main(){
  arr := []int{412,12,123,1,2324,3,155,23412,2}
  fmt.Println(arr)
  arr_order := insertion_sort(arr)
  fmt.Println(arr_order)
}

func insertion_sort(arr []int) []int {
  // starts at 1
  for i := 1; i < len(arr); i++ {
    j := i
    for j > 0 && arr[j-1] > arr[j]{
      swap(j-1, j, &arr) // &arr => mem address of arr
      j--;
    }
  }
  return arr
}

// args as copy
func swap(prev int, actual int, arrp *[]int ){
  // we modify the copy, unless we receive the memory pointer
  arr := *arrp
  copy := arr[actual]
  arr[actual] = arr[prev]
  arr[prev] = copy
}
