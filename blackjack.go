package gotesting

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var validCards = "23456789JQKA"

// CheckIfBlackjack checks if the hand is a blackjack
func CheckIfBlackjack(hand string) bool {
	if len(hand) != 2 {
		return false
	}

	return (hand[0] == 'A' && strings.Contains("JQK", string(hand[1]))) || (hand[1] == 'A' && strings.Contains("JQK", string(hand[0])))
}

// ComputeScore computes the score for a blackjack hand
func ComputeScore(hand string) int {

	total := 0

	for _, card := range hand {
		c := string(card)
		if strings.Contains("JQK", c) {
			total += 10
		} else if c == "A" {
			total += 11
		} else {
			t, err := strconv.Atoi(c)

			if err != nil {
				return 0
			}
			total += t
		}
	}

	return total
}

// CheckIfValidHand checks if the given hand is valid
func CheckIfValidHand(hand string) bool {
	for _, c := range hand {
		if !strings.Contains(validCards, string(c)) {
			return false
		}
	}

	return true
}

// EvaluateHand reads the hand and returns if it's a busted hand, a blackjack orthe score of the hand.
func EvaluateHand(hand string) (string, error) {

	if !CheckIfValidHand(hand) {
		return "", errors.New("Invalid Card found")
	}

	if CheckIfBlackjack(hand) {
		return "BLACKJACK", nil
	}

	score := ComputeScore(hand)

	if score > 22 {
		return "BUSTED", nil
	}

	return strconv.Itoa(score), nil
}

// BlackJackHandler handles requests to check hands
func BlackJackHandler(w http.ResponseWriter, r *http.Request) {

	hand := r.URL.Query().Get("hand")
	if hand == "" {
		w.WriteHeader(422)
		return
	}

	result, err := EvaluateHand(hand)

	if err != nil {
		w.WriteHeader(422)
	}

	w.WriteHeader(200)
	w.Write([]byte(result))

}
