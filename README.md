JSON-RPC client for https://getblock.io.

## Install
```sh
go get github.com/ofen/getblock-go
```

## Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/ofen/getblock-go/eth"
)

func main() {
    ctx := context.Background()

    client := eth.New("api-token")
    head, err := client.BlockNumber(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println(head)
}
```

## Documentation
https://getblock.io/docs/
