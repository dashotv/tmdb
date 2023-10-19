package tmdb

import "testing"

func TestClientWithResponses_Search(t *testing.T) {
	if tmdbURL == "" {
		t.Skip("TMDB_API_URL not set")
	}
	if tmdbToken == "" {
		t.Skip("TMDB_API_TOKEN not set")
	}

}
