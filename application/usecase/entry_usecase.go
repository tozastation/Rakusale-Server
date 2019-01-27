package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pe "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/entry"
	"net/http"
)

// EntryUseCase is ...
type EntryUseCase interface {
	SignIn(ctx context.Context, p *pe.SignInRequest) (*pe.EntryResponse, error)
	SignUp(ctx context.Context, p *pe.SignUpRequest) (*pe.EntryResponse, error)
}

type entryUseCase struct {
	repository.EntryRepository
}

// NewEntryUseCase is ...
func NewEntryUseCase(r repository.EntryRepository) EntryUseCase {
	return &entryUseCase{r}
}

func (u *entryUseCase) SignIn(ctx context.Context, p *pe.SignInRequest) (*pe.EntryResponse, error) {
	res, err := u.EntryRepository.CheckCorrectUser(ctx, p)
	if err != nil {
		res.Status = http.StatusBadRequest
		return res, nil
	}
	res.Status = http.StatusOK
	return res, nil
}

func (u *entryUseCase) SignUp(ctx context.Context, p *pe.SignUpRequest) (*pe.EntryResponse, error) {
	res, err := u.EntryRepository.CreateNewUser(ctx, p)
	if err != nil {
		res.Status = http.StatusBadRequest
		return res, nil
	}
	res.Status = http.StatusOK
	return res, nil
}
