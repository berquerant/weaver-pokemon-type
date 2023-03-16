package persist

//go:generate weaver generate ./...
import (
	"context"
	"database/sql"

	"github.com/ServiceWeaver/weaver"
	"github.com/berquerant/weaver-pokemon-type/domain"
	"github.com/berquerant/weaver-pokemon-type/errorx"
	_ "github.com/go-sql-driver/mysql"
)

type (
	typeDatabaseConfig struct {
		Driver string
		Source string
	}

	TypeDatabase interface {
		GetTypeByName(ctx context.Context, name string) (*domain.Type, error)
	}
	typeDatabase struct {
		weaver.Implements[TypeDatabase]
		weaver.WithConfig[typeDatabaseConfig]

		db *sql.DB
	}
)

func (d *typeDatabase) Init(_ context.Context) error {
	db, err := sql.Open(d.Config().Driver, d.Config().Source)
	d.db = db
	return err
}

func (d *typeDatabase) GetTypeByName(ctx context.Context, name string) (*domain.Type, error) {
	const q = "select id from types where name = ?"
	var id int
	if err := d.db.QueryRowContext(ctx, q, name).Scan(&id); err != nil {
		return nil, errorx.New(err, errorx.WithMessage("GetTypeByName"))
	}
	return &domain.Type{
		ID:   id,
		Name: name,
	}, nil
}
