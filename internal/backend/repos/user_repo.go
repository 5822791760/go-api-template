package repos

import (
	"context"

	"github.com/5822791760/hr/internal/backend/db/hr/public/model"
	t "github.com/5822791760/hr/internal/backend/db/hr/public/table"
	"github.com/5822791760/hr/pkg/coreutil"
	j "github.com/go-jet/jet/v2/postgres"
)

type User struct {
	model.User
	Post []Post
}

type userRepo struct {
	clock coreutil.Clock
}

func NewUserRepo(clock coreutil.Clock) userRepo {
	return userRepo{clock: clock}
}

type UserRepo interface {
	FindOne(ctx context.Context, id int64) (*model.User, error)
	FindAll(ctx context.Context) ([]User, error)
}

func (r userRepo) FindOne(ctx context.Context, id int64) (*model.User, error) {
	var user model.User

	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	qb := j.
		SELECT(t.User.AllColumns).
		FROM(t.User).
		WHERE(
			t.User.ID.EQ(j.Int(id)),
		)

	if err := qb.QueryContext(ctx, db, &user); err != nil {
		if notFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r userRepo) FindAll(ctx context.Context) ([]User, error) {
	var users []User

	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return []User{}, err
	}

	qb := j.
		SELECT(t.User.AllColumns).
		FROM(t.User)

	if err := qb.QueryContext(ctx, db, &users); err != nil {
		return []User{}, err
	}

	return users, nil
}
