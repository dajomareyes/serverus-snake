package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	. "github.com/dajomareyes/serverus-snake/pkg/serverus"
)

var logger *log.Logger

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	blackCape := "#000000"

	response := BattlesnakeInfoResponse {
		APIVersion: "1",
		Author:     "dajomareyes",
		Color:      blackCape,
		Head:       "default", // TODO: Personalize
		Tail:       "default", // TODO: Personalize
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Fatal("Error handling Battlesnake index")
	}
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Fatal("Failed to start battlesnake!")
	}

	// Nothing to respond with here
	logger.Print("START\n")
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	possibleMoves := []string{"up", "down", "left", "right"}
	move := "up"
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	possibleMoves = HandleBoundaries(request, possibleMoves)
	possibleMoves = HandleObstacle(request, possibleMoves)

	if len(possibleMoves) > 0 {
		move = possibleMoves[rand.Intn(len(possibleMoves))]
	} else {
		logger.Println("Death is imminent... goodbye", request.You.Name)
	}

	response := MoveResponse{
		Move: move,
	}

	logger.Println("Next move for", request.You.Name, "is", response)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Println("Failed to send response data")
	}
}

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("END\n")
}

func main() {
	logger = log.New(os.Stdout, "[Main] - ", 5)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)

	logger.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
