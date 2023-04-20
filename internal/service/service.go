package service

import "github.com/dannyhinshaw/pg-patterns/internal/store"

type GibsonHacker interface {
	GetGibson()
	HackGibson()
	StoreGibson()
}

type GibsonHackService struct {
	store store.Postgres
}

func NewGibsonHacker() *GibsonHackService {
	return &GibsonHackService{}
}
