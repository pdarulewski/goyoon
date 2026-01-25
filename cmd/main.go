package main

import (
	"context"
	"fmt"

	"github.com/pdarulewski/goyoon/internal/db"
	"github.com/pdarulewski/goyoon/internal/game"
)

func main() {
	ctx := context.Background()

	pathProvider := &db.XDGStatePathProvider{}
	if err := db.Initialize(pathProvider); err != nil {
		panic(err)
	}

	db, err := db.New(ctx, pathProvider.Path())
	if err != nil {
		panic(err)
	}

	newGame, err := game.New(db)
	if err != nil {
		panic(err)
	}

	gameType := game.FromRomajiToHiragana

	if err := newGame.Prepare(gameType); err != nil {
		panic(err)
	}

	var text string

	for {
		random := newGame.Pick()
		fmt.Printf("%+v\n", random)

		fmt.Print("Enter text: ")
		fmt.Scanln(&text)

		result := newGame.Score(text, random, gameType)
		fmt.Printf("Result: %v\n\n", result.Success)
	}
}
