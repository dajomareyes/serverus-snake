package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func HandleIndex(w http.ResponseWriter, r *http.Request) {

	blackCape := "#000000"

	response := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "dajomareyes", // TODO: Your Battlesnake username
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

func removeMoveOption(possibleMoves *[]string, move string) {
	moves := *possibleMoves
	for i, m := range moves {
		if m == move {
			moves = append(moves[:i], moves[i+1:]...)
		}
	}
	*possibleMoves = moves
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

	// Common variables that will be used throughout the function
	boardWidth := request.Board.Width - 1
	boardHeight := request.Board.Height - 1
	headX := request.You.Head.X
	headY := request.You.Head.Y
	neckX := request.You.Body[1].X
	neckY := request.You.Body[1].Y

	// Choose a random direction to move in
	possibleMoves := []string{"up", "down", "left", "right"}
	arrivedAtWall := headX == 0 || headX == boardWidth || headY == 0 || headY == boardHeight

	move := ""
	// Handle avoiding walls, then figure out how to handle corners
	if arrivedAtWall {
		// choose your next move
		// process of elimination
		// can I move up?
		if headY+1 > boardHeight || headY+1 == neckY {
			removeMoveOption(&possibleMoves, "up")
		}
		// can i move down?
		if headY-1 < 0 || headY-1 == neckY {
			removeMoveOption(&possibleMoves, "down")
		}
		// can i move left?
		if headX-1 < 0 || headX-1 == neckX {
			removeMoveOption(&possibleMoves, "left")
		}
		// can i move right?
		if headX+1 > boardWidth || headX+1 == neckX {
			removeMoveOption(&possibleMoves, "right")
		}
		move = possibleMoves[rand.Intn(len(possibleMoves))]
	}

	response := MoveResponse{
		Move: move,
	}

	// Send response
	fmt.Printf("TURN: %d MOVE: %s POSSIBLE_MOVES: %s NECK: %d,%d HEAD: %d,%d\n", request.Turn, response.Move, possibleMoves, neckX, neckY, headX, headY)
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
