package tmdb

import (
	"context"

	"github.com/dashotv/tmdb/openapi"
)

// Tmdb wrapper client
type Client struct {
	Token string
	sdk   *openapi.SDK
	ctx   context.Context
}

func New(token string) *Client {
	s := openapi.New(openapi.WithSecurity(token))
	return &Client{
		Token: token,
		sdk:   s,
		ctx:   context.Background(),
	}
}
