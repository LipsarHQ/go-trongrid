package trongrid

import (
	"github.com/rs/zerolog"
)

type Option func(api *api)

func WithDebug() Option {
	return func(api *api) {
		api.debug = true
	}
}

func WithLogger(logger *zerolog.Logger) Option {
	return func(api *api) {
		api.logger = logger
	}
}

func WithToken(token string) Option {
	return func(api *api) {
		api.token = token
	}
}

func WithURI(uri string) Option {
	return func(api *api) {
		api.uri = uri
	}
}
