package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	val := 0
	switch card {
	case "ace":
		val = 11
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	case "ten":
		val = 10

	case "jack":
		val = 10
	case "queen":
		val = 10
	case "king":
		val = 10

	default:
		val = 0

	}

	return val
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1Value := ParseCard(card1)
	c2Value := ParseCard(card2)
	cdValue := ParseCard(dealerCard)
	myCardsTotal := c1Value + c2Value
	switch {
	case card1 == card2 && card2 == "ace":
		return "P"
	case myCardsTotal == 21 && cdValue < 10:
		return "W"
	case myCardsTotal == 21 && cdValue >= 10:
		return "S"
	case myCardsTotal >= 17 && myCardsTotal <= 20:
		return "S"
	case myCardsTotal >= 12 && myCardsTotal <= 16 && cdValue < 7:
		return "S"
	case myCardsTotal >= 12 && myCardsTotal <= 16 && cdValue >= 7:
		return "H"
	default:
		return "H"
	}
}
