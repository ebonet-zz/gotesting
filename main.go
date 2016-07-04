package gotesting

import "net/http"

func main() {

	http.HandleFunc("blackjack", BlackJackHandler)
	http.ListenAndServe(":8080", nil)

}

func Sum(x float64, y float64) float64 {
	return x + y
}

func CreateMap(key string, value int) map[string]int {
	m := make(map[string]int)

	m[key] = value

	return m
}
