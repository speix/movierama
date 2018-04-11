package helpers

import "testing"

func TestValidEmail(t *testing.T) {

	email := "spei@supergramm.com"

	if !ValidEmail(email) {
		t.Errorf("Expected %v got %v", true, false)
	}
}

func TestInBetween(t *testing.T) {

	number := 7

	if !InBetween(number, 1, 10) {
		t.Errorf("Expected %v got %v", true, false)
	}

}
