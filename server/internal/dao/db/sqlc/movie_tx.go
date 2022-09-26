package db

import (
	"context"
	"errors"
)

type CreateMovieWithTxParams struct {
	*CreateMovieParams
	Tags []string
}

type CreateMovieWithTx struct {
	*CreateMovieRow
	Tags []string
}

func (store *SqlStore) CreateMovieWithTx(ctx context.Context, params *CreateMovieWithTxParams) (*CreateMovieWithTx, error) {
	ret := new(CreateMovieWithTx)
	err := store.execTx(ctx, func(queries *Queries) (err error) {
		ret.CreateMovieRow, err = queries.CreateMovie(ctx, params.CreateMovieParams)
		if err != nil {
			return err
		}
		tagParams := make([]*CreateTagParams, 0, len(params.Tags))
		for i := range params.Tags {
			tagParams = append(tagParams, &CreateTagParams{
				MovieID: ret.CreateMovieRow.ID,
				TagName: params.Tags[i],
			})
		}
		_, err = queries.CreateTag(ctx, tagParams)
		if err != nil {
			return err
		}
		ret.Tags = params.Tags
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type GetMovieByIDWithTxRow struct {
	*GetMovieByIDRow
	Tags      []string
	IsFollow  bool
	IsComment bool
}

func (store *SqlStore) GetMovieByIDWithTx(ctx context.Context, userID, movieID int64) (ret *GetMovieByIDWithTxRow, err error) {
	ret = new(GetMovieByIDWithTxRow)
	err = store.execTx(ctx, func(queries *Queries) (err error) {
		ret.GetMovieByIDRow, err = queries.GetMovieByID(ctx, movieID)
		if err != nil {
			return err
		}
		ret.Tags, err = queries.GetTagsInMovie(ctx, movieID)
		if err != nil {
			return err
		}
		ret.IsFollow, _ = queries.ExistUserMovie(ctx, &ExistUserMovieParams{
			UserID:  userID,
			MovieID: movieID,
		})
		ret.IsComment, _ = queries.ExistComment(ctx, &ExistCommentParams{
			UserID:  userID,
			MovieID: movieID,
		})
		return nil
	})
	return
}

var ErrMovieHasPlans = errors.New("电影存在演出计划")

func (store *SqlStore) DeleteMovieByIDWithTx(ctx context.Context, movieID int64) error {
	return store.execTx(ctx, func(queries *Queries) error {
		has, err := queries.ExistPlansByMovieID(ctx, movieID)
		if err != nil {
			return err
		}
		if has {
			return ErrMovieHasPlans
		}
		return queries.DeleteMovieByID(ctx, movieID)
	})
}
