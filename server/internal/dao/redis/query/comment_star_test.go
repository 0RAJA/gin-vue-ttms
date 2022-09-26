package query_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/dao/redis/query"
	"ttms/internal/pkg/utils"

	"github.com/stretchr/testify/require"
)

func TestQueries_ListAllCommentStarsAndSetZero(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				require.NoError(t, dao.Group.Redis.ClearCommentStars(context.Background()))
				sumStars := 0
				nums := int(utils.RandomInt(2, 10))
				commentStarMap := make(map[int64]map[int64]bool) // 用户对评论的点赞情况
				for i := 0; i < nums; i++ {
					cnt := int(utils.RandomInt(2, 5))
					sumStars += cnt
					comment := testCreateComment(t)
					for i := 0; i < cnt; i++ {
						user := testCreateUser(t)
						arg := &query.ListCommentStarsReply{
							UserID:    user.ID,
							CommentID: comment.ID,
							State:     rand.Intn(2) == 0,
						}
						if commentStarMap[user.ID] == nil {
							commentStarMap[user.ID] = make(map[int64]bool)
						}
						require.NoError(t, dao.Group.Redis.SetCommentStar(context.Background(), arg.UserID, arg.CommentID, arg.State))
						commentStarMap[user.ID][comment.ID] = arg.State
					}
				}
				now := time.Now()
				result, err := dao.Group.Redis.GetHalfCommentStarsAndSetZero(context.Background())
				require.NoError(t, err)
				t.Log(time.Since(now))
				t.Log(len(result), sumStars)
				require.True(t, len(result) >= sumStars/2 && len(result) < sumStars)
				for i := range result {
					require.EqualValues(t, commentStarMap[result[i].UserID][result[i].CommentID], result[i].State)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}

func testCreateUser(t *testing.T) *db.CreateUserRow {
	pwd, err := utils.RandomPassword(10)
	require.NoError(t, err)
	user, err := dao.Group.DB.CreateUser(context.Background(), &db.CreateUserParams{
		Username:  utils.RandomOwner(),
		Password:  pwd,
		Avatar:    utils.RandomAvatar(),
		Email:     utils.RandomEmail(),
		Birthday:  time.Now(),
		Privilege: db.PrivilegeValue2,
		Signature: utils.RandomString(20),
	})
	require.NoError(t, err)
	return user
}

func testCreateMovie(t *testing.T) *db.CreateMovieWithTx {
	tagNum := int(utils.RandomInt(1, 10))
	tags := make([]string, 0, tagNum)
	tagMap := map[string]bool{}
	for i := 0; i < tagNum; i++ {
		tag := utils.RandomTag()
		if tagMap[tag] == false {
			tags = append(tags, tag)
			tagMap[tag] = true
		}
	}
	movie, err := dao.Group.DB.CreateMovieWithTx(context.Background(), &db.CreateMovieWithTxParams{
		CreateMovieParams: &db.CreateMovieParams{
			Name:      utils.RandomOwner(),
			AliasName: utils.RandomOwner(),
			Director:  utils.RandomOwner(),
			Actors:    utils.RandomStringSlice(10, 10),
			Content:   utils.RandomString(10),
			Avatar:    utils.RandomAvatar(),
			Duration:  int16(utils.RandomInt(1, 100)),
			Area:      utils.RandomArea(),
			Period:    utils.RandomPeriod(),
		},
		Tags: tags,
	})
	require.NoError(t, err)
	return movie
}

func testCreateComment(t *testing.T) *db.Comment {
	user := testCreateUser(t)
	movie := testCreateMovie(t)
	comment, err := dao.Group.DB.CreateComment(context.Background(), &db.CreateCommentParams{
		Content:   utils.RandomString(20),
		MovieID:   movie.ID,
		UserID:    user.ID,
		Score:     float32(utils.RandomFloat(1, 10)),
		IpAddress: "127.0.0.1",
	})
	require.NoError(t, err)
	return comment
}

func TestQueries_SetCommentStar(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				comment := testCreateComment(t)
				err := dao.Group.Redis.SetCommentStar(context.Background(), comment.UserID, comment.ID, true)
				require.NoError(t, err)
				result, err := dao.Group.Redis.GetCommentStar(context.Background(), comment.UserID, comment.ID)
				require.NoError(t, err)
				require.Equal(t, result, "true")
				err = dao.Group.Redis.SetCommentStar(context.Background(), comment.UserID, comment.ID, false)
				require.NoError(t, err)
				result, err = dao.Group.Redis.GetCommentStar(context.Background(), comment.UserID, comment.ID)
				require.NoError(t, err)
				require.Equal(t, result, "false")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}
