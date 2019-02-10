package usecase

import (
	"context"
	"fmt"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
	"github.com/sirupsen/logrus"
	"net/http"
)

// VegetableUseCase is ...
type VegetableUseCase interface {
	GetMyBoughtVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error)
	GetMySoldVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error)
	GetAllVegetables(ctx context.Context, p *pv.VegetablesEmpty) (*pv.GetAllVegetablesResponse, error)
	GetSingleShopAllVegetables(ctx context.Context, p *pv.GetSingleShopAllVegetablesRequest) (*pv.GetSingleShopAllVegetablesResponse, error)
	PostMyVegetable(ctx context.Context, p *pv.PostMyVegetableRequest) (*pv.PostMyVegetableResponse, error)
	PutMyVegetable(ctx context.Context, p *pv.PutMyVegetableRequest) (*pv.PutMyVegetableResponse, error)
	DeleteMyVegetable(ctx context.Context, p *pv.DeleteMyVegetableRequest) (*pv.DeleteMyVegetableResponse, error)
	BuyVegetables(ctx context.Context, p *pv.BuyVegetablesRequest) (*pv.BuyVegetablesResponse, error)
}

type vegetableUseCase struct {
	repository.VegetableRepository
	Logger *logrus.Logger
}

// NewVegetableUseCase is ...
func NewVegetableUseCase(r repository.VegetableRepository, l *logrus.Logger) VegetableUseCase {
	return &vegetableUseCase{r, l}
}

func (u *vegetableUseCase) GetMyBoughtVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error) {
	u.Logger.Infoln("[START] GetMyBoughtVegetablesRPC is Called from Client")
	res := pv.GetMyVegetablesResponse{}
	token := p.GetToken()
	u.Logger.Infoln("[INPUT] token=" + token)
	vegetables, err := u.VegetableRepository.FindMyBoughtVegetables(ctx, token)
	if err != nil {
		u.Logger.Error("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusNoContent
		return &res, nil
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	u.Logger.Infoln("[END] GetMyBoughtVegetablesRPC is Called from Client")
	return &res, nil
}

func (u *vegetableUseCase) GetMySoldVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error) {
	u.Logger.Infoln("[START] GetMySoldVegetablesRPC is Called from Client")
	res := pv.GetMyVegetablesResponse{}
	token := p.GetToken()
	u.Logger.Infoln("[INPUT] token=" + token)
	vegetables, err := u.VegetableRepository.FindMySoldVegetables(ctx, token)
	if err != nil {
		u.Logger.Error("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusNoContent
		return &res, nil
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	u.Logger.Infoln("[END] GetMySoldVegetablesRPC is Called from Client")
	return &res, nil
}

func (u *vegetableUseCase) GetAllVegetables(ctx context.Context, p *pv.VegetablesEmpty) (*pv.GetAllVegetablesResponse, error) {
	u.Logger.Infoln("[START] GetAllVegetablesRPC is Called from Client")
	res := pv.GetAllVegetablesResponse{}
	vegetables, err := u.VegetableRepository.FindAllVegetables(ctx)
	if err != nil {
		u.Logger.Error("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusNoContent
		return &res, nil
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	u.Logger.Infoln("[END] GetAllVegetablesRPC is Called from Client")
	return &res, nil
}

func (u *vegetableUseCase) GetSingleShopAllVegetables(ctx context.Context, p *pv.GetSingleShopAllVegetablesRequest) (*pv.GetSingleShopAllVegetablesResponse, error) {
	u.Logger.Infoln("[START] GetSingleShopAllVegetablesRPC is Called from Client")
	res := pv.GetSingleShopAllVegetablesResponse{}
	sID := p.GetShopId()
	u.Logger.Info("[INPUT] shopID=%s", sID)
	vegetables, err := u.VegetableRepository.FindSingleShopAllVegetables(ctx, sID)
	if err != nil {
		u.Logger.Debug("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusNoContent
		return &res, nil
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	u.Logger.Infoln("[END] GetSingleShopAllVegetablesRPC is Called from Client")
	return &res, nil
}

func (u *vegetableUseCase) PostMyVegetable(ctx context.Context, p *pv.PostMyVegetableRequest) (*pv.PostMyVegetableResponse, error) {
	u.Logger.Infoln("[START] PostMyVegetableRPC is Called from Client")
	res := pv.PostMyVegetableResponse{}
	token := p.GetToken()
	u.Logger.Infoln("[INPUT] token=" + token)
	err := u.VegetableRepository.AddMyVegetable(ctx, token, p)
	if err != nil {
		u.Logger.Error("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusBadRequest
		return &res, nil
	}
	res.Status = http.StatusCreated
	u.Logger.Infoln("[END] PostMyVegetableRPC is Called from Client")
	return &res, nil
}

func (u *vegetableUseCase) PutMyVegetable(ctx context.Context, p *pv.PutMyVegetableRequest) (*pv.PutMyVegetableResponse, error) {
	u.Logger.Infoln("[START] PutMyVegetableRPC is Called from Client")
	res := pv.PutMyVegetableResponse{}
	err := u.VegetableRepository.UpdateMyVegetable(ctx, p.GetToken(), p.GetVId(), p.GetVegetable())
	if err != nil {
		u.Logger.Error("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusBadRequest
		return &res, nil
	}
	res.Status = http.StatusCreated
	u.Logger.Infoln("[END] PutMyVegetableRPC is Called from Client")
	return &res, nil
}

func (u *vegetableUseCase) DeleteMyVegetable(ctx context.Context, p *pv.DeleteMyVegetableRequest) (*pv.DeleteMyVegetableResponse, error) {
	u.Logger.Infoln("[START] DeleteMyVegetableRPC is Called from Client")
	res := pv.DeleteMyVegetableResponse{}
	err := u.VegetableRepository.DeleteMyVegetable(ctx, p.GetToken(), p.GetVId())
	if err != nil {
		u.Logger.Error("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusBadRequest
		return &res, nil
	}
	res.Status = http.StatusAccepted
	u.Logger.Infoln("[END] DeleteMyVegetableRPC is Called from Client")
	return &res, nil
}

func (u *vegetableUseCase) BuyVegetables(ctx context.Context, p *pv.BuyVegetablesRequest) (*pv.BuyVegetablesResponse, error) {
	u.Logger.Infoln("[START] BuyVegetablesRPC is Called from Client")
	res := pv.BuyVegetablesResponse{}
	err := u.VegetableRepository.BuyVegetables(ctx, p.GetToken(), p.GetSID(), fmt.Sprint(p.GetCategory()), p.GetAmount())
	if err != nil {
		u.Logger.Error("[EXECUTE FAILURE!] %s\n", err)
		res.Status = http.StatusBadRequest
		return &res, nil
	}
	res.Status = http.StatusCreated
	return &res, nil
}
