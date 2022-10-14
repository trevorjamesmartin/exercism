package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var dealer = ParseCard(dealerCard)
	var hand = ParseCard(card1) + ParseCard(card2)
	switch {
	case hand == 22: // split two aces.
		return "P"
	case dealer >= 10 && hand == 21: // blackjack, wait for the reveal
		return "S"
	case dealer < 10 && hand == 21: // blackjack, for the win
		return "W"
	case hand < 21 && hand > 16:
		return "S"
	case hand < 17 && hand > 11 && dealer > 6:
		return "H"
	case hand < 17 && hand > 11:
		return "S"
	default:
		return "H"
	}
}
