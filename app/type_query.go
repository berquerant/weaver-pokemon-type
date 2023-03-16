package app

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/berquerant/weaver-pokemon-type/domain"
	"github.com/berquerant/weaver-pokemon-type/persist"
	"github.com/berquerant/weaver-pokemon-type/weaverx"
)

type (
	GetTypeByNameQuery interface {
		GetTypeByName(ctx context.Context, name string) (*Type, error)
	}

	getTypeByNameQuery struct {
		weaver.Implements[GetTypeByNameQuery]

		tdb persist.TypeDatabase
	}
)

func (q *getTypeByNameQuery) Init(_ context.Context) error {
	tdb, err := weaver.Get[persist.TypeDatabase](q)
	q.tdb = tdb
	return err
}

func (q *getTypeByNameQuery) GetTypeByName(ctx context.Context, name string) (*Type, error) {
	item, err := weaverx.Retry(func() (*domain.Type, error) {
		return q.tdb.GetTypeByName(ctx, name)
	})
	if err != nil {
		return nil, err
	}
	return &Type{
		ID:   item.ID,
		Name: item.Name,
	}, nil
}
