package helpers

import "testing"

func TestCorrectPasswordHash(t *testing.T) {

	expected := "gne3F6R2MLmSAXEyRtUCakaNCYXFghufe2Y9W1MVdCLmIPs6lYvn3TdMFbsmg-WbuZTTOzsZEPFm5yWc2vewLQ=="
	email := "spei@supergramm.com"
	pass := "123456789"

	hashed := HashPassword(email, pass)

	if expected != hashed {
		t.Errorf("Expected %v got %v", expected, hashed)
	}

}
