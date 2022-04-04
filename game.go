package main

import "fmt"

//Declare Struct for Game
type game struct {
	title string
	price float64
}

//Declare method for Game
func (game game) printPrice() {
	fmt.Printf("Title: %v  Price: $%.2f\n", game.title, game.price)
}