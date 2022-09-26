package logic

import (
	"context"
	"errors"

	"ttms/internal/dao"
	db "ttms/internal/dao/db/sqlc"
	"ttms/internal/global"
	"ttms/internal/pkg/goroutine/task"
	"ttms/internal/pkg/utils"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

/*
	自动化任务
*/
type auto struct {
}

func (auto) Work() {
	ctx := context.Background()
	movieVisitCountFlushTask := task.Task{
		Name:            "movieVisitCountFlushTask",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.MovieVisitCountDuration,
		TimeoutDuration: global.Settings.Server.DefaultContextTimeout,
		F:               movieVisitCountFlush(),
	}
	commentStarFlushTask := task.Task{
		Name:            "commentStarFlushTask",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.CommentStarDuration,
		TimeoutDuration: global.Settings.Server.DefaultContextTimeout,
		F:               commentStarFlush(),
	}
	setMoviesOrderByBoxOfficeTask := task.Task{
		Name:            "setMoviesOrderByBoxOfficeTask",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.MoviesOrderByBoxOfficeDuration,
		TimeoutDuration: global.Settings.Server.DefaultContextTimeout,
		F:               setMoviesOrderByBoxOffice(),
	}
	setMoviesOrderByUserMovieCountTask := task.Task{
		Name:            "setMoviesOrderByUserMovieCountTask",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.MoviesOrderByUserMovieCountDuration,
		TimeoutDuration: global.Settings.Server.DefaultContextTimeout,
		F:               setMoviesOrderByUserMovieCount(),
	}
	deleteOutDatePlansTask := task.Task{
		Name:            "deleteOutDatePlansTask",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.DeleteOutDatePlansDuration,
		TimeoutDuration: global.Settings.Server.DefaultContextTimeout,
		F:               deleteOutDatePlans(),
	}
	addVisitCountTask := task.Task{
		Name:            "addVisitCount",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.AddVisitCountDuration,
		TimeoutDuration: global.Settings.Server.DefaultContextTimeout,
		F:               addVisitCount(),
	}
	startTask(movieVisitCountFlushTask, commentStarFlushTask, setMoviesOrderByBoxOfficeTask, setMoviesOrderByUserMovieCountTask, deleteOutDatePlansTask, addVisitCountTask)
}

func startTask(tasks ...task.Task) {
	for i := range tasks {
		task.NewTickerTask(tasks[i])
	}
}

// 自动刷新电影访问量
func movieVisitCountFlush() task.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("auto task run: movieVisitCountFlush")
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Server.DefaultContextTimeout)
		defer cancel()
		movieVisitCountMap, err := dao.Group.Redis.GetAllMovieVisitCountAndSetZero(ctx)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		for id, count := range movieVisitCountMap {
			if err := dao.Group.DB.AddMovieVisitCount(ctx, &db.AddMovieVisitCountParams{
				ID:         id,
				VisitCount: count,
			}); err != nil {
				global.Logger.Error(err.Error())
			}
		}
	}
}

// 刷新评论点赞情况
func commentStarFlush() task.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("auto task run: commentStarFlush")
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Server.DefaultContextTimeout)
		defer cancel()
		commentStars, err := dao.Group.Redis.GetHalfCommentStarsAndSetZero(ctx)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		for _, commentStar := range commentStars {
			if err := optCommentStar(ctx, commentStar.UserID, commentStar.CommentID, commentStar.State); err != nil {
				if err.Error() != "ERROR: duplicate key value violates unique constraint \"comment_star_pkey\" (SQLSTATE 23505)" {
					global.Logger.Error(err.Error(), zap.Int64("userID", commentStar.UserID), zap.Int64("commentID", commentStar.CommentID))
				}
			}
		}
	}
}

// 根据commentStar的opt对DB进行操作
func optCommentStar(ctx context.Context, userID, commentID int64, opt bool) error {
	if opt {
		return dao.Group.DB.CreateCommentStar(ctx, &db.CreateCommentStarParams{UserID: userID, CommentID: commentID})
	}
	return dao.Group.DB.DeleteCommentStar(ctx, &db.DeleteCommentStarParams{UserID: userID, CommentID: commentID})
}

// DeleteOutDatePlans 自动删除过期演出计划
func deleteOutDatePlans() task.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("auto task run: deleteOutDatePlans")
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Server.DefaultContextTimeout)
		defer cancel()
		_, err := dao.Group.DB.DeleteOutDatePlans(ctx)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
	}
}

// setMoviesOrderByBoxOffice 更新电影票房排序到缓存
func setMoviesOrderByBoxOffice() task.DoFunc {
	return func(parentCtx context.Context) {
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Server.DefaultContextTimeout)
		defer cancel()
		page, size := global.Settings.Rule.MoviesOrderByBoxOfficePage, global.Settings.Rule.MoviesOrderByBoxOfficeSize
		data, err := dao.Group.DB.GetMoviesOrderByBoxOffice(ctx, page*size)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		var start, end int
		for i := 1; i <= int(page); i++ {
			start = (i - 1) * int(size)
			end = start + int(size)
			if end > len(data) {
				end = len(data)
			}
			if start >= end {
				break
			}
			if err := dao.Group.Redis.Set(ctx, utils.LinkStr(MoviesOrderByBoxOfficeKey, utils.IDToSting(int64(i))), data[start:end]); err != nil {
				global.Logger.Error(err.Error())
				return
			}
		}
	}
}

// setMoviesOrderByUserMovieCount 更新电影根据期待值排序到缓存
func setMoviesOrderByUserMovieCount() task.DoFunc {
	return func(parentCtx context.Context) {
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Server.DefaultContextTimeout)
		defer cancel()
		page, size := global.Settings.Rule.MoviesOrderByUserMovieCountPage, global.Settings.Rule.MoviesOrderByUserMovieCountSize
		data, err := dao.Group.DB.GetMoviesOrderByUserMovieCount(ctx, page*size)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		var start, end int
		for i := 1; i <= int(page); i++ {
			start = (i - 1) * int(size)
			end = start + int(size)
			if end > len(data) {
				end = len(data)
			}
			if start >= end {
				break
			}
			if err := dao.Group.Redis.Set(ctx, utils.LinkStr(MoviesOrderByUserMovieCountKey, utils.IDToSting(int64(i))), data[start:end]); err != nil {
				global.Logger.Error(err.Error())
				return
			}
		}
	}
}

// 刷新访问量到数据库
func addVisitCount() task.DoFunc {
	return func(parentCtx context.Context) {
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Server.DefaultContextTimeout)
		defer cancel()
		num, err := dao.Group.Redis.GetVisitNumAndSetZero(ctx)
		if err != nil && !errors.Is(err, redis.Nil) {
			global.Logger.Error(err.Error())
			return
		}
		if num == 0 {
			return
		}
		if err := dao.Group.DB.CreateVisitCount(ctx, num); err != nil {
			global.Logger.Error(err.Error())
		}
	}
}
