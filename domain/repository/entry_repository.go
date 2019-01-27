package repository

import (
	"context"
	pe "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/entry"
)

// EntryRepository is Interface of Infrastructure Repository
type EntryRepository interface {
	CheckCorrectUser(ctx context.Context, p *pe.SignInRequest) (*pe.EntryResponse, error)
	CreateNewUser(ctx context.Context, p *pe.SignUpRequest) (*pe.EntryResponse, error)
}
