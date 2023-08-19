package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "benchung.dev", nil)
	req.Header.Set("Authorization", "ApiKey benchhung")
	_, err := GetAPIKey(req.Header)
	if err != nil {
		fmt.Printf("%s", err)
		t.Fatal("Error!")
	}
}
