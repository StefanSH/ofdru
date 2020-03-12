package ofdru

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/Authorization/CreateAuthToken" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{"AuthToken": "f3accdfda7574736ba94a78d00e974f4","ExpirationDateUtc": "2017-01-24T14:13:24"}`))
			}
		}),
	)
	defer ts.Close()

	client := OfdRu("", "12345", "56789", ts.URL)

	at, err := client.auth()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, at.AuthToken, "f3accdfda7574736ba94a78d00e974f4")
	assert.Equal(t, at.ExpirationDateUtc, time.Date(2017,01,24,14,13,24,0, time.UTC))
}
