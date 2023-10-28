<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	"github.com/dashotv/tmdb/openapi"
	"github.com/dashotv/tmdb/openapi/models/operations"
	"github.com/dashotv/tmdb/openapi/models/shared"
	"log"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity(""),
	)

	var accountID int = 959345

	requestBody := &operations.AccountAddFavoriteRequestBody{
		RawBody: "{key: 16540, key1: null, key2: \"string\"}",
	}

	var sessionID *string = "string"

	ctx := context.Background()
	res, err := s.SDK.AccountAddFavorite(ctx, accountID, requestBody, sessionID)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountAddFavorite200ApplicationJSONObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->