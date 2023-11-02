package trongrid

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

type API interface {
	// ListTransactions
	// Docs: https://developers.tron.network/reference/get-trc20-transaction-info-by-account-address
	ListTransactions(
		ctx context.Context,
		req *ListTransactionsRequest,
	) (resp *ListTransactionsResponse, err error)
}

type api struct {
	encoder *schema.Encoder
	decoder *schema.Decoder
	logger  *zerolog.Logger
	cl      *resty.Client
	token   string
	uri     string
	debug   bool
}

func NewAPI(opts ...Option) API {
	x := &api{
		encoder: NewEncoder(),
		decoder: NewDecoder(),
		logger:  nil,
		cl:      nil,
		token:   "",
		uri:     "",
		debug:   false,
	}
	for _, opt := range opts {
		opt(x)
	}

	if len(x.uri) == 0 {
		x.uri = URI
	}

	cl := resty.New().
		SetBaseURL(x.uri).
		SetDebug(x.debug).
		SetRateLimiter(rate.NewLimiter(rate.Every(time.Second), 1)).
		SetRedirectPolicy(resty.NoRedirectPolicy()).
		SetTimeout(timeout)
	if x.logger != nil {
		cl.SetLogger(NewLogger(x.logger))
	}

	if len(x.token) != 0 {
		cl.SetHeader("TRON-PRO-API-KEY", x.token)
	}

	x.cl = cl

	return x
}
