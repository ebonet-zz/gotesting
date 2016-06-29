package go_testing_presentation

import (
	"strings"
	"strconv"
	"errors"
	"net/http"
)

var validCards = "23456789JQKA"

func CheckIfBlackjack(hand string) bool {
	if len(hand) != 2 {
		return false
	}

	return (hand[0]=='A' && strings.Contains("JQK", string(hand[1]))) || (hand[1]=='A' && strings.Contains("JQK", string(hand[0])))
}

func ComputeScore(hand string) int {

	total := 0

	for _, card := range hand {
		c := string(card)
		if strings.Contains("JQK", c) {
			total +=10
		} else if c =="A" {
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

func CheckIfValid(hand string) bool{
	for _, c := range hand {
		if !strings.Contains(validCards, string(c)) {
			return false
		}
	}

	return true
}

func CheckHand(hand string) (string, error) {

	if !CheckIfValid(hand) {
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

func BlackJackHandler(w http.ResponseWriter, r *http.Request) {

	hand := r.URL.Query().Get("hand")
	if hand == "" {
		w.WriteHeader(422)
		return
	}

	result, err := CheckHand(hand)

	if err != nil {
		w.WriteHeader(422)
	}

	w.WriteHeader(200)
	w.Write([]byte(result))

}

func main() {
	http.HandleFunc("/", BlackJackHandler)
	http.ListenAndServe(":8080", nil)
}


