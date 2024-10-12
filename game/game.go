package game

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Game struct {
	Players []Player
	Movies  []Movie
	Rules   []Rules
}

type Rules struct {
	NumberOfPlayers int
	StrikeOuts      int
	Votes           int
	MoviesToAdd     int
}

type Player struct {
	ID   uuid.UUID
	Name string
}

type Movie struct {
	ID    uuid.UUID
	Title string
	Votes int
	// ToDo: Add ImdbRating
}

func NewGame() *Game {
	return &Game{}
}

func StartGame() {
	g := NewGame()
	g.explainRules()
	g.setupGameRules()
	g.addPlayers()
}

func (g *Game) explainRules() {
	fmt.Println("The rules of movie list are as follows:")
	fmt.Println("1. During the first rounds each randomly selected player may add a chosen amount of movies.")
	fmt.Println("2. During the second round each player may strike out a selected amount of movies.")
	fmt.Println("3. During the third round the players watch trailers for all still available movies.")
	fmt.Println("4. During the fourth and final round the players each cast their vote for their movie of choice.")
	fmt.Println("5. The movie with the most votes wins.")
	fmt.Println("6. In case of a tie the players may vote again.")
	fmt.Println()
}

func (g *Game) setupGameRules() {
	var rules Rules

	fmt.Println("How many players are playing?")
	_, err := fmt.Scan(&rules.NumberOfPlayers)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("How many movies can each player add?")
	_, err = fmt.Scan(&rules.MoviesToAdd)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("How many strike outs are allowed?")
	_, err = fmt.Scan(&rules.StrikeOuts)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("How many votes are allowed?")
	_, err = fmt.Scan(&rules.Votes)
	if err != nil {
		log.Fatal(err)
		return
	}

	g.Rules = []Rules{rules}
}

func (g *Game) addPlayers() {
	for i := 0; i < g.Rules[0].NumberOfPlayers; i++ {
		id := uuid.New()
		fmt.Println("Enter player name:")
		var name string
		_, err := fmt.Scan(&name)
		if err != nil {
			log.Fatal(err)
			return
		}
		g.Players = append(g.Players, Player{ID: id, Name: name})
	}

	g.printPlayers()
}

func (g *Game) printPlayers() {
	fmt.Println("The players are:")
	for i := range g.Players {
		fmt.Println(g.Players[i].Name)
	}
}
