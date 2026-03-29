package blackjack

const (
    Stand	= "S"
    Hit 	= "H"
    Split	= "P"
    Win		= "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
        case "ace":
        	return 11
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
        case "ten", "jack", "queen", "king":
        	return 10
        default:
        	return 0
    }
    return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var sum = ParseCard(card1) + ParseCard(card2)
    var dealerCardVal = ParseCard(dealerCard)

    switch {
		// If you have a pair of aces you must always split them.
        case card1 == card2 && card1 == "ace":
        	return Split
    	// If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a face card (Jack/Queen/King) or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
        case sum == 21 && dealerCardVal < 10:
        	return Win
        case sum == 21 && dealerCardVal >= 10:
        	return Stand
		// If your cards sum up to a value within the range [17, 20] you should always stand.
        case sum >= 17 && sum <= 20:
        	return Stand
		// If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
        case sum >= 12 && sum <= 16 && dealerCardVal < 7:
        	return Stand
        case sum >= 12 && sum <= 16 && dealerCardVal >= 7:
        	return Hit
		// If your cards sum up to 11 or lower you should always hit.
        case sum <= 11:
        	return Hit
    }

    return ""
}
