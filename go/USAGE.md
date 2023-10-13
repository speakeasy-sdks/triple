<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/triple"
	"log"
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
<!-- End SDK Example Usage -->