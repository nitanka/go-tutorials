package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

var randomJokes = []string{
	"Why don't scientists trust atoms? Because they make up everything!",
	"Why did the scarecrow win an award? Because he was outstanding in his field!",
	"Why don't skeletons fight each other? They don't have the guts!",
	"Why did the bicycle fall over? Because it was two-tired!",
	"Why did the tomato turn red? Because it saw the salad dressing!",
}

func getRandomJoke(jokes []string) string {
	random := rand.Intn(len(randomJokes))
	return jokes[random]
}

func returnJokeHandler(w http.ResponseWriter, r *http.Request) {
	joke := getRandomJoke(randomJokes)
	fmt.Fprintf(w, joke)
}

func main() {
	http.HandleFunc("/what-a-joke", returnJokeHandler)
	fmt.Println("Starting the server at 8080")
	http.ListenAndServe(":8080", nil)
}
