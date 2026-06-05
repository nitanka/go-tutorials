package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubbleSortHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var numbers []int
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()

	err := decoder.Decode(&numbers)

	fmt.Println("Received request to sort numbers ", r.Body)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	bubbleSort(numbers)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(numbers)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {

	http.HandleFunc("/sort", bubbleSortHandler)
	fmt.Println("Starting the server at 8080")
	http.ListenAndServe(":8080", nil)
}

/*

curl -X POST http://localhost:8080/sort \
  -H "Content-Type: application/json" \
  -d "[5, 3, 8, 1, 4]"
[1,3,4,5,8]

*/
