package sorting

func SelectionSort(array []int) {
	n := len(array)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
}

// func main() {
// 	var (
// 		size int
// 		numbers []int
// 	)
// 	fmt.Println("Enter the number of elements")
// 	fmt.Scanln(&size)

// 	fmt.Println("Please enter the elements")
// 	for i := 0; i<size; i++ {
// 		var element int
// 		fmt.Scanln(&element)
// 		numbers = append(numbers, element)
// 	}
// 	fmt.Println("The Array before merging ", numbers)
// 	selectionSort(numbers)
// 	fmt.Println("The Arrar after merging ", numbers)

// }
