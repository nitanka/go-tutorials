package sorting

import "fmt"

func merging(array []int, low int, mid int, high int) {
	fmt.Println("\nMerging ", array)
	temp := make([]int, high-low+1)
	var (
		l1 int = low
		l2 int = mid + 1
		i  int = 0
	)
	for l1 <= mid && l2 <= high {
		if array[l1] <= array[l2] {
			temp[i] = array[l1]
			l1++
		} else {
			temp[i] = array[l2]
			l2++
		}
		i++
	}
	for l1 <= mid {
		temp[i] = array[l1]
		l1++
		i++
	}
	for l2 <= high {
		temp[i] = array[l2]
		l2++
		i++
	}

	for k := 0; k < len(temp); k++ {
		array[low+k] = temp[k]
	}
	fmt.Println("\nMerged array: ", array)
}

func MergeSort(array []int, low int, high int) {
	fmt.Printf("\n\nSorting %v, low: %d, high: %d\n", array, low, high)
	var mid int
	if low < high {
		mid = (low + high) / 2
		MergeSort(array, low, mid)
		MergeSort(array, mid+1, high)
		merging(array, low, mid, high)
	} else {
		fmt.Printf("\n\nBase case reached for array: %v, low: %d, high: %d\n", array, low, high)
		return
	}
}

// func main() {
//    var array []int
//    fmt.Println("Enter the number of elements in the array")
//    var n int
//    fmt.Scanln(&n)
//    fmt.Println("Enter the elements of the array")
//    for i := 0; i < n; i++ {
// 	  var element int
// 	  fmt.Scanln(&element)
// 	  array = append(array, element)
//    }
//    fmt.Println("Given array is", array)
//    sorting(array, 0, len(array)-1)
//    fmt.Println("Sorted array is", array)
// }
