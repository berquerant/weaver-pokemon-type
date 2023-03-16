package persist

import (
	"context"
	"database/sql"

	"github.com/ServiceWeaver/weaver"
	"github.com/berquerant/weaver-pokemon-type/domain"
	"github.com/berquerant/weaver-pokemon-type/errorx"
	_ "github.com/go-sql-driver/mysql"
)

type (
	effectivityDatabaseConfig struct {
		Driver string
		Source string
	}

	EffectivityDatabase interface {
		GetEffectivityByAttack(ctx context.Context, attackTypeID int) ([]*domain.Effectivity, error)
		GetEffectivityByDefense(ctx context.Context, defenseTypeID int) ([]*domain.Effectivity, error)
	}
	effectivityDatabase struct {
		weaver.Implements[EffectivityDatabase]
		weaver.WithConfig[effectivityDatabaseConfig]

		db *sql.DB
	}
)

func (d *effectivityDatabase) Init(_ context.Context) error {
	db, err := sql.Open(d.Config().Driver, d.Config().Source)
	d.db = db
	return err
}

type Effectivity struct {
	ID          int     `db:"id"`
	AttackID    int     `db:"attack_id"`
	AttackName  string  `db:"attack_name"`
	DefenseID   int     `db:"defense_id"`
	DefenseName string  `db:"defense_name"`
	Multiplier  float32 `db:"multiplier"`
}

func (d *effectivityDatabase) GetEffectivityByAttack(ctx context.Context, attackTypeID int) ([]*domain.Effectivity, error) {
	const q = `select e.id as id, a.id as attack_id, a.name as attack_name, d.id as defense_id, d.name as defense_name, e.multiplier as multiplier
from effectivities e
inner join types a on a.id = e.attack_type_id
inner join types d on d.id = e.defense_type_id
where e.attack_type_id = ?`
	load := func(rows *sql.Rows) (*domain.Effectivity, error) {
		var (
			id          int
			attackID    int
			attackName  string
			defenseID   int
			defenseName string
			multiplier  float32
		)
		if err := rows.Scan(&id, &attackID, &attackName, &defenseID, &defenseName, &multiplier); err != nil {
			return nil, err
		}
		return &domain.Effectivity{
			ID: id,
			Attack: domain.Type{
				ID:   attackID,
				Name: attackName,
			},
			Defense: domain.Type{
				ID:   defenseID,
				Name: defenseName,
			},
			Multiplier: multiplier,
		}, nil
	}

	rows, err := d.db.QueryContext(ctx, q, attackTypeID)
	if err != nil {
		return nil, errorx.New(err, errorx.WithMessage("GetEffectivityByAttack"))
	}
	defer rows.Close()

	list := make([]*domain.Effectivity, 0)
	for rows.Next() {
		item, err := load(rows)
		if err != nil {
			return nil, errorx.New(err, errorx.WithMessage("GetEffectivityByAttack"))
		}
		list = append(list, item)
	}
	if err := rows.Err(); err != nil {
		return nil, errorx.New(err, errorx.WithMessage("GetEffectivityByAttack"))
	}
	return list, nil
}

func (d *effectivityDatabase) GetEffectivityByDefense(ctx context.Context, defenseTypeID int) ([]*domain.Effectivity, error) {
	const q = `select e.id as id, a.id as attack_id, a.name as attack_name, d.id as defense_id, d.name as defense_name, e.multiplier as multiplier
from effectivities e
inner join types a on a.id = e.attack_type_id
inner join types d on d.id = e.defense_type_id
where e.defense_type_id = ?`
	load := func(rows *sql.Rows) (*domain.Effectivity, error) {
		var (
			id          int
			attackID    int
			attackName  string
			defenseID   int
			defenseName string
			multiplier  float32
		)
		if err := rows.Scan(&id, &attackID, &attackName, &defenseID, &defenseName, &multiplier); err != nil {
			return nil, err
		}
		return &domain.Effectivity{
			ID: id,
			Attack: domain.Type{
				ID:   attackID,
				Name: attackName,
			},
			Defense: domain.Type{
				ID:   defenseID,
				Name: defenseName,
			},
			Multiplier: multiplier,
		}, nil
	}

	rows, err := d.db.QueryContext(ctx, q, defenseTypeID)
	if err != nil {
		return nil, errorx.New(err, errorx.WithMessage("GetEffectivityByDefense"))
	}
	defer rows.Close()

	list := make([]*domain.Effectivity, 0)
	for rows.Next() {
		item, err := load(rows)
		if err != nil {
			return nil, errorx.New(err, errorx.WithMessage("GetEffectivityByDefense"))
		}
		list = append(list, item)
	}
	if err := rows.Err(); err != nil {
		return nil, errorx.New(err, errorx.WithMessage("GetEffectivityByDefense"))
	}
	return list, nil
}
