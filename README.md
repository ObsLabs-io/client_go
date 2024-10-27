# ObsLabs API Client Go

## Installation

```sh
go get github.com/obslabs-io/client_go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/ObsLabs-io/client_go"
)

func main() {
	ctx := context.WithValue(context.Background(), client_go.ContextBasicAuth, client_go.BasicAuth{
		Password: "API_KEY",
	})
	client := client_go.NewAPIClient(client_go.NewConfiguration())

	organizationID := "org_123"
	projectID := "pro_abc"
	probeID := "prb_456"

	resp, _, err := client.ProbesAPI.V1GetProbe(ctx, organizationID, projectID, probeID).Execute()
	if err != nil {
		panic(err)
	}

	probe := resp.GetProbe()

	fmt.Printf("Probe status: %v\n", probe.GetStatus())
}
```
