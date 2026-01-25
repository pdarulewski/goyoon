package game

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/pdarulewski/goyoon/internal/db"
	"github.com/pdarulewski/goyoon/pkg/kana"
)

type Type int

const (
	FromHiraganaToRomaji Type = iota
	FromRomajiToHiragana

	FromKatakanaToRomaji
	FromRomajiToKatakana

	FromRomajiToKana
)

type Game struct {
	DB   *db.DB
	Type Type

	kana       kana.Kana
	dictionary kana.Dictionary
}

func New(db *db.DB) (*Game, error) {
	if db == nil {
		return nil, fmt.Errorf("")
	}

	return &Game{
		DB: db,
	}, nil
}

func (g *Game) Prepare(gameType Type) error {
	switch gameType {
	case FromRomajiToHiragana:
		if err := g.prepareGame(kana.Hiragana, kana.FromRomajiToKana); err != nil {
			return fmt.Errorf("")
		}

	case FromRomajiToKatakana:
		if err := g.prepareGame(kana.Katakana, kana.FromRomajiToKana); err != nil {
			return fmt.Errorf("")
		}

	case FromRomajiToKana:
		bothKanas := make(kana.Kana, len(kana.Hiragana)+len(kana.Katakana))
		_ = copy(bothKanas, kana.Hiragana)
		_ = copy(bothKanas[len(bothKanas):], kana.Katakana)

		if err := g.prepareGame(bothKanas, kana.FromRomajiToKana); err != nil {
			return fmt.Errorf("")
		}

	case FromHiraganaToRomaji:
		if err := g.prepareGame(kana.Hiragana, kana.FromKanaToRomaji); err != nil {
			return fmt.Errorf("")
		}

	case FromKatakanaToRomaji:
		if err := g.prepareGame(kana.Katakana, kana.FromKanaToRomaji); err != nil {
			return fmt.Errorf("")
		}

	default:
		return fmt.Errorf("")
	}

	return nil
}

func (g *Game) Pick() kana.Mora {
	return g.kana[rand.Intn(len(g.kana))]
}

type Result struct {
	Success bool
}

func (g *Game) Score(guess string, target kana.Mora, gameType Type) Result {
	log.Printf("guess: `%s`", guess)
	v, ok := g.dictionary[guess]
	if !ok {
		return Result{
			Success: false,
		}
	}

	log.Printf("guess: %s, target: %+v, equal: %v", v, target, v.Character == target.Character)

	switch gameType {
	case FromHiraganaToRomaji, FromKatakanaToRomaji:
		return Result{
			Success: target.Romaji == v.Romaji,
		}

	case FromRomajiToHiragana, FromRomajiToKatakana:
		return Result{
			Success: target.Character == v.Character,
		}

	case FromRomajiToKana:
		return Result{
			Success: true,
		}
	}

	return Result{
		Success: false,
	}
}

func (g *Game) prepareGame(kana kana.Kana, direction kana.Direction) error {
	g.kana = kana

	dictionary, err := g.kana.ToDictionary(direction)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	g.dictionary = dictionary

	return nil
}
