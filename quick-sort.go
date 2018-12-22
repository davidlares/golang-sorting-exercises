package main

import "fmt"

func main(){
  arr := []int {31,41,59,26,53,58,97,93,23,84,40,1}
  quickSort(&arr,0,len(arr)-1)
  fmt.Println(arr)
}

func quickSort(arr *[]int, left,right int){
  // spliting values (left from pivot) and bigger values to the right
  if left < right {
      a := *arr
      limit, origin := right,left
      pivot :=  a[right]
      right--

      for left <= right {
        // search from left side a bigger number than pivot
        for left < len(a) && a[left] < pivot {
          left++
        }
        // search from right side a smaller number than pivot
        for right >= 0 && a[right] > pivot {
          right--
        }
        // flip
        if left <= right {
          swap(arr, left, right)
          left++
          right--
        }
      }
      // end separation
      swap(arr,left,limit)
      quickSort(arr,origin,right)
      quickSort(arr,left,limit)

  }
}

func swap(arr *[]int, left, right int){
  a := *arr
  temp := a[left]
  a[left] = a[right]
  a[right] = temp
}
