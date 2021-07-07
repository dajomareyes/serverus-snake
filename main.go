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

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	blackCape := "#000000"

	response := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "dajomareyes",
		Color:      blackCape,
		Head:       "default", // TODO: Personalize
		Tail:       "default", // TODO: Personalize
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
// TODO: Use this function to decide how your Battlesnake is going to look on the board.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("START\n")
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
// TODO: Use the information in the GameRequest object to determine your next move.
func HandleMove(w http.ResponseWriter, r *http.Request) {
	// Receive request
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	move := "up" // we default to up if we can't select a move
	// possible moves that you can select
	possibleMoves := []string{"up", "down", "left", "right"}

	/*
		TODO: this process allows for us to change direction when we are at a wall
		the problem right now is that if we're not moving towards the wall but running along side it
		this will cause the snake to turn too. This is a side affect of this logic.

		HandleBoundaries & HandleObstacle both take a possible list of moves and updates the moves by removing
		the ones that cannot be used
	*/
	possibleMoves = HandleBoundaries(request, possibleMoves)
	possibleMoves = HandleObstacle(request, possibleMoves)

	if len(possibleMoves) > 0 {
		move = possibleMoves[rand.Intn(len(possibleMoves))]
	} else {
		fmt.Println("You're trapped... goodbye")
	}

	response := MoveResponse{
		Move: move,
	}

	fmt.Println(request.You.ID, response)

	// Send response
	// fmt.Printf("TURN: %d MOVE: %s POSSIBLE_MOVES: %s NECK: %d,%d HEAD: %d,%d\n", request.Turn, response.Move, possibleMoves, neckX, neckY, headX, headY)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
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
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
