package trongrid_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/LipsarHQ/go-trongrid"
)

func TestApi_ListTransactions(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	api := trongrid.NewAPI(
		trongrid.WithDebug(),
		trongrid.WithLogger(&logger),
		trongrid.WithToken(""),
	)

	ctx := context.Background()
	now := time.Now()

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
	require.NoError(t, err)
	require.NotNil(t, modelListTransactionsRequest)
}
