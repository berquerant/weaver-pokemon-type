package service

//go:generate weaver generate ./...
import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/berquerant/weaver-pokemon-type/app"
	"github.com/berquerant/weaver-pokemon-type/weaverx"
)

type (
	GetTypeByNameRequest struct {
		weaver.AutoMarshal
		Name string `json:"name"`
	}
	GetTypeByNameResponse struct {
		weaver.AutoMarshal
		Item *app.Type `json:"item"`
	}
)

type (
	API interface {
		GetTypeByName(ctx context.Context, request *GetTypeByNameRequest) (*GetTypeByNameResponse, error)
		GetEffectivityListByAttack(ctx context.Context, request *GetEffectivityListByAttackRequest) (*GetEffectivityListByAttackResponse, error)
		GetEffectivityListByDefenseList(ctx context.Context, request *GetEffectivityListByDefenseListRequest) (*GetEffectivityListByDefenseListResponse, error)
	}

	api struct {
		weaver.Implements[API]

		getTypeByName                   app.GetTypeByNameQuery
		getEffectivityListByAttack      app.GetEffectivityListByAttackQuery
		getEffectivityListByDefenseList app.GetEffectivityListByDefenseListQuery
	}
)

func (a *api) Init(_ context.Context) error {
	getEffectivityListByDefenseList, err := weaver.Get[app.GetEffectivityListByDefenseListQuery](a)
	if err != nil {
		return err
	}
	a.getEffectivityListByDefenseList = getEffectivityListByDefenseList
	getEffectivityListByAttack, err := weaver.Get[app.GetEffectivityListByAttackQuery](a)
	if err != nil {
		return err
	}
	a.getEffectivityListByAttack = getEffectivityListByAttack
	getTypeByName, err := weaver.Get[app.GetTypeByNameQuery](a)
	a.getTypeByName = getTypeByName
	return err
}

func (a *api) GetTypeByName(ctx context.Context, request *GetTypeByNameRequest) (*GetTypeByNameResponse, error) {
	result, err := weaverx.Retry(func() (*app.Type, error) {
		return a.getTypeByName.GetTypeByName(ctx, request.Name)
	})
	if err != nil {
		return nil, err
	}
	return &GetTypeByNameResponse{
		Item: result,
	}, nil
}

type (
	GetEffectivityListByAttackRequest struct {
		weaver.AutoMarshal
		ID    int           `json:"id"`
		Index app.PileIndex `json:"index"`
	}
	GetEffectivityListByAttackResponseItem struct {
		weaver.AutoMarshal
		Pile       app.PiledEffectivity `json:"pile"`
		Multiplier float32              `json:"multiplier"`
	}
	GetEffectivityListByAttackResponse struct {
		weaver.AutoMarshal
		Items []GetEffectivityListByAttackResponseItem `json:"items"`
	}
)

func (a *api) GetEffectivityListByAttack(ctx context.Context, request *GetEffectivityListByAttackRequest) (*GetEffectivityListByAttackResponse, error) {
	result, err := weaverx.Retry(func() ([]app.PiledEffectivity, error) {
		return a.getEffectivityListByAttack.GetEffectivityListByAttack(ctx, request.Index, request.ID)
	})
	if err != nil {
		return nil, err
	}
	items := make([]GetEffectivityListByAttackResponseItem, len(result))
	for i, r := range result {
		items[i] = GetEffectivityListByAttackResponseItem{
			Pile:       r,
			Multiplier: r.Multiplier(),
		}
	}
	return &GetEffectivityListByAttackResponse{
		Items: items,
	}, nil
}

type (
	GetEffectivityListByDefenseListRequest struct {
		weaver.AutoMarshal
		DefenseTypeIDs app.DefenseTypeIDList `json:"ids"`
	}
	GetEffectivityListByDefenseListResponseItem struct {
		weaver.AutoMarshal
		Pile       app.PiledEffectivity `json:"pile"`
		Multiplier float32              `json:"multiplier"`
	}
	GetEffectivityListByDefenseListResponse struct {
		weaver.AutoMarshal
		Items []GetEffectivityListByDefenseListResponseItem `json:"items"`
	}
)

func (a *api) GetEffectivityListByDefenseList(ctx context.Context, request *GetEffectivityListByDefenseListRequest) (*GetEffectivityListByDefenseListResponse, error) {
	result, err := weaverx.Retry(func() ([]app.PiledEffectivity, error) {
		return a.getEffectivityListByDefenseList.GetEffectivityListByDefenseList(ctx, request.DefenseTypeIDs)
	})
	if err != nil {
		return nil, err
	}
	items := make([]GetEffectivityListByDefenseListResponseItem, len(result))
	for i, r := range result {
		items[i] = GetEffectivityListByDefenseListResponseItem{
			Pile:       r,
			Multiplier: r.Multiplier(),
		}
	}
	return &GetEffectivityListByDefenseListResponse{
		Items: items,
	}, nil
}
