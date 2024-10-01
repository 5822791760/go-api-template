package repos

import (
	"context"

	"github.com/5822791760/hr/internal/backend/db/hr/public/model"
	t "github.com/5822791760/hr/internal/backend/db/hr/public/table"
	"github.com/5822791760/hr/pkg/coreutil"
	j "github.com/go-jet/jet/v2/postgres"
)

type Post struct {
	model.Post
	User *User
}

type postRepo struct {
	clock coreutil.Clock
}

func NewPostRepo(clock coreutil.Clock) postRepo {
	return postRepo{clock: clock}
}

type PostRepo interface {
	FindOne(ctx context.Context, id int) (*Post, error)
	FindAll(ctx context.Context) ([]Post, error)
}

func (r postRepo) FindOne(ctx context.Context, id int64) (*Post, error) {
	var post Post

	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	qb := j.
		SELECT(
			t.Post.AllColumns,
		).
		FROM(t.Post).
		WHERE(
			t.Post.ID.EQ(j.Int(id)),
		)

	if err := qb.QueryContext(ctx, db, &post); err != nil {
		if notFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return &post, nil
}

func (r postRepo) FindAll(ctx context.Context) ([]Post, error) {
	var posts []Post

	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return []Post{}, err
	}

	qb := j.SELECT(
		t.Post.AllColumns,
	).
		FROM(t.Post)

	if err := qb.QueryContext(ctx, db, &posts); err != nil {
		return []Post{}, err
	}

	return posts, nil
}
