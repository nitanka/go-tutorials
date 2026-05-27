package main

import "fmt"


func merging(array []int, low int, mid int, high int) {
   fmt.Println("Merging", array)
   var (
      l1 int = low
      l2 int = mid + 1
      i  int = low
   )
   for l1 <= mid && l2 <= high {
      if array[l1] <= array[l2] {
         array[i] = array[l1]
         l1++
      } else {
         array[i] = array[l2]
         l2++
      }
      i++
   }
   for l1 <= mid {
      array[i] = array[l1]
      l1++
      i++
   }
   for l2 <= high {
      array[i] = array[l2]
      l2++
      i++
   }
}

func sorting(array []int, low int, high int) {
   var mid int;
   if(low < high) {
      mid = (low + high) / 2;
      sorting(array, low, mid);
      sorting(array, mid+1, high);
      merging(array, low, mid, high);
   } else {
      return;
   }
}


func main() {
   var array []int
   fmt.Println("Enter the number of elements in the array")
   var n int
   fmt.Scanln(&n)
   fmt.Println("Enter the elements of the array")
   for i := 0; i < n; i++ {
	  var element int
	  fmt.Scanln(&element)
	  array = append(array, element)
   }
   fmt.Println("Given array is", array)
   sorting(array, 0, len(array)-1)
   fmt.Println("Sorted array is", array)
}