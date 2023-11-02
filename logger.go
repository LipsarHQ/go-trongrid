package trongrid

import (
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
)

type logger struct {
	l *zerolog.Logger
}

func NewLogger(l *zerolog.Logger) resty.Logger {
	return &logger{l: l}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.l.Debug().Msgf(format, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.l.Error().Msgf(format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.l.Warn().Msgf(format, v...)
}
