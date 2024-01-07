<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	triple "github.com/speakeasy-sdks/triple/v3"
	"log"
	"net/http"
)

func main() {
	s := triple.New()

	ctx := context.Background()
	res, err := s.Pets.CreatePets(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->