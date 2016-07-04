package gotesting

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestComputeScore(t *testing.T) {
	if ComputeScore("2") != 2 {
		t.Error("Score of 2 should be 2")
	}

	if ComputeScore("3") != 3 {
		t.Error("Score of 3 should be 3")
	}

	if ComputeScore("J") != 10 {
		t.Error("Score of J should be 10")
	}

	if ComputeScore("A") != 11 {
		t.Error("Score of A should be 11")
	}

	if ComputeScore("AJ") != 21 {
		t.Error("Score of AJ should be 21")
	}

	if ComputeScore("2245") != 13 {
		t.Error("Score of 2245 should be 13")
	}

	if ComputeScore("A345") != 23 {
		t.Error("Score of A345 should be 23")
	}

	if ComputeScore("/345") != 0 {
		t.Error("Score of invalid should be 0")
	}

	if ComputeScore("") != 0 {
		t.Error("Score of empty should be 0")
	}
}

func TestComputeScore_Better(t *testing.T) {

	testCases := []struct {
		Hand          string
		ExpectedScore int
	}{
		{"2", 2},
		{"3", 3},
		{"J", 10},
		{"A", 11},
		{"AJ", 21},
		{"A345", 23},
		{"/345≈", 0}, // Invalid character
		{"", 0},      // Empty hand
	}

	for _, tc := range testCases {
		if ComputeScore(tc.Hand) != tc.ExpectedScore {
			t.Errorf("Score for %s should be %d",
				tc.Hand, tc.ExpectedScore)
		}
	}
}

func TestCheckIfBlackjack(t *testing.T) {

	testCases := []struct {
		Hand           string
		ExpectedResult bool
	}{
		{"AJ", true},
		{"JA", true},
		{"A4", false},
		{"J2", false},
		{"2", false},
		{"AJ5", false},
	}

	for _, tc := range testCases {
		if CheckIfBlackjack(tc.Hand) != tc.ExpectedResult {
			t.Errorf("%s should be %d", tc.Hand, tc.ExpectedResult)
		}
	}
}

func TestCheckIfValidHand(t *testing.T) {
	testCases := []struct {
		Hand           string
		ExpectedResult bool
	}{
		{"AJ", true},
		{"AAAA", true},
		{"12345", false},
		{"//1234", false},
	}

	for _, tc := range testCases {
		if CheckIfValidHand(tc.Hand) != tc.ExpectedResult {
			t.Errorf("%s should be %d", tc.Hand, tc.ExpectedResult)
		}
	}
}

func TestEvaluateHand(t *testing.T) {
	testCases := []struct {
		Hand           string
		ExpectedResult string
		ExpectErr      bool
	}{
		{"AJ", "BLACKJACK", false},
		{"AAAA", "BUSTED", false},
		{"12345", "", true},
		{"//1234", "", true},
		{"A3", "14", false},
		{"3235", "13", false},
	}

	for _, tc := range testCases {

		result, err := EvaluateHand(tc.Hand)

		if tc.ExpectErr {
			if err == nil {
				t.Errorf("Case %s should err, but did not", tc.Hand)
			}
		} else if err != nil {
			t.Errorf("Case %s should not err, but did %s", tc.Hand, err)
		} else {
			if result != tc.ExpectedResult {
				t.Errorf("Case %s returned %s, but was expecting %s",
					tc.Hand, result, tc.ExpectedResult)
			}
		}
	}
}

func TestBlackJackHandler(t *testing.T) {

	req, _ := http.NewRequest("GET", "http://test.com/", nil)
	recorder := httptest.NewRecorder()
	BlackJackHandler(recorder, req)

	if recorder.Code != 422 {
		t.Errorf("Wrong response to no hand in query ")
	}

	req.URL.RawQuery = url.Values{"hand": []string{"AJ"}}.Encode()

	recorder = httptest.NewRecorder()
	BlackJackHandler(recorder, req)

	if recorder.Code != 200 {
		t.Errorf("Wrong status code for AJ: Expected 200, got %d",
			recorder.Code)
	} else if b := recorder.Body.String(); b != "BLACKJACK" {
		t.Errorf("Wrong response body for AJ: Expected BLACKJACK, got %s",
			b)
	}

}
