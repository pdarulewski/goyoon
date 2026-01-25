// Package kana contains the definition of the syllabaries.
package kana

import (
	"errors"
	"fmt"
)

// ErrKana is a general kana error.
var ErrKana = errors.New("kana error")

// Type defines a type of character.
type Type string

// Types of modifiers.
const (
	Basic      Type = "basic"
	Dakuten    Type = "dakuten"
	Handakuten Type = "handakuten"
	Yoon       Type = "yoon"
)

// Mora represents a single character.
type Mora struct {
	Character string
	Romaji    string
	Type      Type
}

// Kana is a list of characters.
type Kana []Mora

// Direction defines an order of translation used to create a map.
type Direction int

// Directions for conversion a list of kana to a map.
const (
	FromKanaToRomaji Direction = iota
	FromRomajiToKana
)

// Dictionary represents kana as a map.
type Dictionary map[string]string

// ToDictionary converts a list of kana to a map according to given direction.
func (k *Kana) ToDictionary(direction Direction) (Dictionary, error) {
	dictionary := make(Dictionary, len(*k))

	switch direction {
	case FromKanaToRomaji:
		for _, v := range *k {
			dictionary[v.Character] = v.Romaji
		}
	case FromRomajiToKana:
		for _, v := range *k {
			dictionary[v.Romaji] = v.Character
		}
	default:
		return nil, fmt.Errorf("%w: wrong direction", ErrKana)
	}

	return dictionary, nil
}
