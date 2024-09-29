package repos

import (
	"context"

	"github.com/5822791760/hr/pkg/coreutil"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/sqlscan"
)

var UserT = "user"

type userRepo struct {
	clock coreutil.Clock
}

func NewUserRepo(clock coreutil.Clock) userRepo {
	return userRepo{clock: clock}
}

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type UserRepo interface {
	FindOne(ctx context.Context, id int) (*User, error)
}

func (r userRepo) FindOne(ctx context.Context, id int) (*User, error) {
	var user User

	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return &User{}, err
	}

	sql, _, err := goqu.From(UserT).
		Select(&user).
		Where(goqu.C("id").Eq(id)).
		ToSQL()

	if err != nil {
		return &User{}, err
	}

	if err := sqlscan.Get(ctx, db, &user, sql); err != nil {
		if sqlscan.NotFound(err) {
			return nil, nil
		}

		return &User{}, err
	}

	return &user, nil
}
