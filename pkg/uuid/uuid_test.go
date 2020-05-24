package uuid

import "testing"

func Test_GenerateUUIDSuccess(t *testing.T) {
	u := Generate()
	if u == "" {
		t.Error("UUID not expected to be empty string")
	}
}

func Test_generateFallbackUUIDSuccess(t *testing.T) {
	u := generateFallback()
	if u == "" {
		t.Error("UUID not expected to be empty string")
	}
	if len(u) != 36 {
		t.Errorf("Expected lenght of 36 but recieved %d", len(u))
	}
	// t.Log(u)
}
