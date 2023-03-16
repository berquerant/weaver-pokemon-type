package app

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/berquerant/weaver-pokemon-type/domain"
	"github.com/berquerant/weaver-pokemon-type/errorx"
	"github.com/berquerant/weaver-pokemon-type/persist"
	"github.com/berquerant/weaver-pokemon-type/weaverx"
	combinations "github.com/mxschmitt/golang-combinations"
)

type (
	GetEffectivityListByAttackQuery interface {
		GetEffectivityListByAttack(ctx context.Context, index PileIndex, attackTypeID int) ([]PiledEffectivity, error)
	}

	getEffectivityListByAttackQuery struct {
		weaver.Implements[GetEffectivityListByAttackQuery]

		edb persist.EffectivityDatabase
	}
)

func (q *getEffectivityListByAttackQuery) Init(_ context.Context) error {
	edb, err := weaver.Get[persist.EffectivityDatabase](q)
	q.edb = edb
	return err
}

func (q *getEffectivityListByAttackQuery) GetEffectivityListByAttack(ctx context.Context, index PileIndex, attackTypeID int) ([]PiledEffectivity, error) {
	if err := index.Validate(); err != nil {
		return nil, errorx.New(err, errorx.WithStatusCode(errorx.BadRequest))
	}

	effectivities, err := weaverx.Retry(func() ([]*domain.Effectivity, error) {
		return q.edb.GetEffectivityByAttack(ctx, attackTypeID)
	})
	if err != nil {
		return nil, err
	}

	var result []PiledEffectivity
	for _, xs := range combinations.Combinations(effectivities, int(index)) {
		var p PiledEffectivity
		for _, x := range xs {
			p = p.Append(Effectivity{
				ID: x.ID,
				Attack: Type{
					ID:   x.Attack.ID,
					Name: x.Attack.Name,
				},
				Defense: Type{
					ID:   x.Defense.ID,
					Name: x.Defense.Name,
				},
				Multiplier: x.Multiplier,
			})
		}
		result = append(result, p)
	}
	return result, nil
}

type (
	GetEffectivityListByDefenseListQuery interface {
		GetEffectivityListByDefenseList(ctx context.Context, defenseTypeIDList DefenseTypeIDList) ([]PiledEffectivity, error)
	}

	getEffectivityListByDefenseListQuery struct {
		weaver.Implements[GetEffectivityListByDefenseListQuery]

		edb persist.EffectivityDatabase
	}
)

func (q *getEffectivityListByDefenseListQuery) Init(_ context.Context) error {
	edb, err := weaver.Get[persist.EffectivityDatabase](q)
	q.edb = edb
	return err
}

func (q *getEffectivityListByDefenseListQuery) GetEffectivityListByDefenseList(ctx context.Context, defenseTypeIDList DefenseTypeIDList) ([]PiledEffectivity, error) {
	if err := defenseTypeIDList.Validate(); err != nil {
		return nil, errorx.New(err, errorx.WithStatusCode(errorx.BadRequest))
	}

	list := make(map[int]PiledEffectivity) // attack_type_id -> PiledEffectivity
	for _, id := range defenseTypeIDList {
		xs, err := weaverx.Retry(func() ([]*domain.Effectivity, error) {
			return q.edb.GetEffectivityByDefense(ctx, id)
		})
		if err != nil {
			return nil, err
		}

		for _, x := range xs {
			e := list[x.Attack.ID]
			list[x.Attack.ID] = e.Append(Effectivity{
				ID: x.ID,
				Attack: Type{
					ID:   x.Attack.ID,
					Name: x.Attack.Name,
				},
				Defense: Type{
					ID:   x.Defense.ID,
					Name: x.Defense.Name,
				},
				Multiplier: x.Multiplier,
			})
		}
	}

	var (
		result = make([]PiledEffectivity, len(list))
		i      int
	)
	for _, xs := range list {
		result[i] = xs
		i++
	}
	return result, nil
}
