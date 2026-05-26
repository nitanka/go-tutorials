package main

import "fmt" 


func bubblesort(numberList []int) []int {
	n := len(numberList)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numberList[j] > numberList[j+1] {
				temp := numberList[j]
				numberList[j] = numberList[j+1]
				numberList[j+1] = temp
			}
		}
	}
	return numberList

}

func main() {

	var numbers [5]int //list to store the number 
	fmt.Println("Hello we are doing a bubblesort")

	fmt.Println("Please enter a list of 5 numbers")

	for i := 0; i < 5; i++ {
		fmt.Println("Enter the number")
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Printing the list of numbers")
	for i := 0; i < 5; i++ {
		fmt.Printf("The element in %d => %d\n", i, numbers[i])
	}

	sortedList := bubblesort(numbers[:])



	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 5-i-1; j++ {
	// 		if numberList[j] > numberList[j+1] {
	// 			temp := numberList[j]
	// 			numberList[j] = numberList[j+1]
	// 			numberList[j+1] = temp
	// 		}
	// 	}
	// }
	fmt.Println("The original list is ", numbers)

	fmt.Println("The sorted list is ", sortedList)

}