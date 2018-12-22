package main

import "fmt"

func main(){
  original := []int{4,412,12,123,1,2324,3,155,2,23412}
  fmt.Println(original)
  ordered := bubble_sort(original)
  fmt.Println(ordered)
}

func bubble_sort(arr []int) []int {
  ordered_num := 0
  iteration_num := 0
  flip := true
  for flip {  // flip until flip is true
    flip = false
    for i := 1; (i < len(arr) - ordered_num); i++ {
      iteration_num++
      // compare
      if arr[i - 1] > arr[i]{
        // [10,3] -> flip
        flip = true
        swap(&arr, i) // memory address
      }
    }
    ordered_num++
  }
  fmt.Println(iteration_num)
  return arr
}

func swap(arr_pointer *[]int, rightindex int){
  leftindex := rightindex - 1
  arr := *arr_pointer // returns value on arr_pointer (whole array)
  copy := arr[leftindex]
  arr[leftindex] = arr[rightindex]
  arr[rightindex] = copy
}
