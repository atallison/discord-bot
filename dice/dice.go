package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

type DiceRoll struct {
	Dice []Die
}

type Die struct {
	Sides int
	Value int
}

func Roll(diceString string) (*DiceRoll, error) {

	dice, err := parseSingleDie(diceString)
	if err != nil {
		return nil, err
	}

	return &DiceRoll{Dice: dice}, nil
}

var ErrInvalidDiceString = fmt.Errorf("Invalid dice string")

func parseSingleDie(diceString string) ([]Die, error) {
	exp, err := regexp.Compile("(\\d+)d(\\d+)")
	if err != nil {
		return nil, err
	}
	matches := exp.FindStringSubmatch(diceString)
	if err != nil {
		return nil, err
	}
	numOfDice, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}
	diceValue, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	dice := make([]Die, numOfDice)
	for i := 0; i < numOfDice; i++ {
		dice[i] = Die{
			Sides: diceValue,
			Value: random(diceValue) + 1,
		}
	}
	return dice, nil
}

func random(max int) int {
	return rand.Intn(max)
}
