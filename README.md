# TronGrid API

:star: Star us on GitHub — it motivates us a lot!

Go/Golang library for the TronGrid API.

## Table of Content

- [Features](#features)
- [Usage/Examples](#usageexamples)
    - [List Transactions](#list-transactions)
- [License](#license)
- [Links](#links)

## Features

- List Transactions

## Usage/Examples

#### List Transactions

```go
package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/LipsarHQ/go-trongrid"
)

const token = "13f70f8e-65ff-49d2-a218-7cf3fadf6230"

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	api := trongrid.NewAPI(
		trongrid.WithDebug(),
		trongrid.WithLogger(&logger),
		trongrid.WithToken(token),
	)

	ctx := context.Background()
	now := time.Now()

	// List transactions.
	modelListTransactionsRequest, err := api.ListTransactions(ctx, &trongrid.ListTransactionsRequest{
		MaxTimestamp:  now.Add(-(time.Hour * 24)),
		MinTimestamp:  now,
		Address:       "TWpMnUh9pZS1Mf8yyw9WPiS82WYevKzQo2",
		Fingerprint:   "",
		OrderBy:       "block_timestamp,desc",
		Limit:         2,
		OnlyConfirmed: true,
		OnlyFrom:      false,
		OnlyTo:        false,
	})
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Links

* [TronGrid API](https://developers.tron.network/reference/background)
