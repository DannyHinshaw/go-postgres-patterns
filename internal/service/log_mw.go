package service

import (
	"golang.org/x/exp/slog"
)

type GibsonHackerLogger struct {
	log  *slog.Logger
	next GibsonHacker
}

func NewServiceLogger(s GibsonHacker, log *slog.Logger) GibsonHackerLogger {
	return GibsonHackerLogger{
		log:  log,
		next: s,
	}
}
