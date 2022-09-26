// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddMovieBoxOffice(ctx context.Context, arg *AddMovieBoxOfficeParams) error
	AddMovieVisitCount(ctx context.Context, arg *AddMovieVisitCountParams) error
	CheckCinemaByName(ctx context.Context, name string) (bool, error)
	CheckUserRepeat(ctx context.Context, arg *CheckUserRepeatParams) (int64, error)
	CreateCinema(ctx context.Context, arg *CreateCinemaParams) (*Cinema, error)
	CreateComment(ctx context.Context, arg *CreateCommentParams) (*Comment, error)
	CreateCommentStar(ctx context.Context, arg *CreateCommentStarParams) error
	CreateMovie(ctx context.Context, arg *CreateMovieParams) (*CreateMovieRow, error)
	CreateOrder(ctx context.Context, arg *CreateOrderParams) error
	CreatePlan(ctx context.Context, arg *CreatePlanParams) (int64, error)
	CreateSeat(ctx context.Context, arg *CreateSeatParams) error
	CreateSeats(ctx context.Context, arg []*CreateSeatsParams) (int64, error)
	CreateTag(ctx context.Context, arg []*CreateTagParams) (int64, error)
	CreateTicket(ctx context.Context, arg *CreateTicketParams) error
	CreateTickets(ctx context.Context, arg []*CreateTicketsParams) (int64, error)
	CreateUser(ctx context.Context, arg *CreateUserParams) (*CreateUserRow, error)
	CreateUserMovie(ctx context.Context, arg *CreateUserMovieParams) error
	CreateVisitCount(ctx context.Context, visitCount int64) error
	DeleteByMovieId(ctx context.Context, movieID int64) error
	DeleteByPlan(ctx context.Context, planid int64) error
	DeleteBySeats(ctx context.Context, arg *DeleteBySeatsParams) error
	DeleteByTagName(ctx context.Context, tagName string) error
	DeleteCinemaByID(ctx context.Context, id int64) error
	DeleteCommentByID(ctx context.Context, id int64) error
	DeleteCommentStar(ctx context.Context, arg *DeleteCommentStarParams) error
	DeleteMovieByID(ctx context.Context, id int64) error
	DeleteOneByMovieAndTag(ctx context.Context, arg *DeleteOneByMovieAndTagParams) error
	DeleteOrderByTicket(ctx context.Context, arg *DeleteOrderByTicketParams) error
	DeleteOrderByUUID(ctx context.Context, orderID uuid.UUID) error
	DeleteOutDatePlans(ctx context.Context) ([]int64, error)
	DeletePlan(ctx context.Context, id int64) error
	DeleteSeat(ctx context.Context, arg *DeleteSeatParams) error
	DeleteSeatsByPlan(ctx context.Context, cinemaid int64) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteUserById(ctx context.Context, id int64) error
	DeleteUserMovie(ctx context.Context, arg *DeleteUserMovieParams) error
	ExistComment(ctx context.Context, arg *ExistCommentParams) (bool, error)
	ExistPlansByCinemaID(ctx context.Context, cinemaID int64) (bool, error)
	ExistPlansByMovieID(ctx context.Context, movieID int64) (bool, error)
	// 查询演出计划是否有已经售出或者锁定的票
	ExistSoldTicketsByPlan(ctx context.Context, planid int64) (bool, error)
	ExistUserMovie(ctx context.Context, arg *ExistUserMovieParams) (bool, error)
	GetAllPlanIds(ctx context.Context) ([]*GetAllPlanIdsRow, error)
	GetAllTickets(ctx context.Context, arg *GetAllTicketsParams) ([]*GetAllTicketsRow, error)
	GetAreas(ctx context.Context, arg *GetAreasParams) ([]string, error)
	GetByPlan(ctx context.Context, planid int64) ([]*GetByPlanRow, error)
	GetCinemaByID(ctx context.Context, id int64) (*Cinema, error)
	GetCinemaByPlanID(ctx context.Context, id int64) (*Cinema, error)
	GetCinemas(ctx context.Context, arg *GetCinemasParams) ([]*GetCinemasRow, error)
	GetCommentByID(ctx context.Context, id int64) (*GetCommentByIDRow, error)
	GetCommentStar(ctx context.Context, arg *GetCommentStarParams) (*CommentStar, error)
	GetCommentsByMovieID(ctx context.Context, arg *GetCommentsByMovieIDParams) ([]*GetCommentsByMovieIDRow, error)
	GetCommentsByUserID(ctx context.Context, arg *GetCommentsByUserIDParams) ([]*GetCommentsByUserIDRow, error)
	GetMovieByID(ctx context.Context, movieID int64) (*GetMovieByIDRow, error)
	GetMovieInTag(ctx context.Context, tagName string) ([]*GetMovieInTagRow, error)
	GetMovies(ctx context.Context, arg *GetMoviesParams) ([]*GetMoviesRow, error)
	GetMoviesByIDs(ctx context.Context, ids []int64) ([]*GetMoviesByIDsRow, error)
	GetMoviesByNameOrContent(ctx context.Context, arg *GetMoviesByNameOrContentParams) ([]*GetMoviesByNameOrContentRow, error)
	GetMoviesByTagPeriodAreaOrderByPeriod(ctx context.Context, arg *GetMoviesByTagPeriodAreaOrderByPeriodParams) ([]*GetMoviesByTagPeriodAreaOrderByPeriodRow, error)
	GetMoviesByTagPeriodAreaOrderByScore(ctx context.Context, arg *GetMoviesByTagPeriodAreaOrderByScoreParams) ([]*GetMoviesByTagPeriodAreaOrderByScoreRow, error)
	GetMoviesByTagPeriodAreaOrderByVisitCount(ctx context.Context, arg *GetMoviesByTagPeriodAreaOrderByVisitCountParams) ([]*GetMoviesByTagPeriodAreaOrderByVisitCountRow, error)
	GetMoviesOrderByBoxOffice(ctx context.Context, limit int32) ([]*GetMoviesOrderByBoxOfficeRow, error)
	GetMoviesOrderByUserMovieCount(ctx context.Context, limit int32) ([]*GetMoviesOrderByUserMovieCountRow, error)
	GetMoviesOrderByVisitCount(ctx context.Context, arg *GetMoviesOrderByVisitCountParams) ([]*GetMoviesOrderByVisitCountRow, error)
	GetNumsAll(ctx context.Context) (int64, error)
	GetOne(ctx context.Context, arg *GetOneParams) (*GetOneRow, error)
	GetOrderByUserId(ctx context.Context, userID int64) ([]*GetOrderByUserIdRow, error)
	GetOrderInfoByCinemaId(ctx context.Context, arg *GetOrderInfoByCinemaIdParams) (*GetOrderInfoByCinemaIdRow, error)
	GetPlanByID(ctx context.Context, id int64) (*GetPlanByIDRow, error)
	GetPlans(ctx context.Context, arg *GetPlansParams) ([]*GetPlansRow, error)
	GetPlansByMovie(ctx context.Context, arg *GetPlansByMovieParams) ([]*GetPlansByMovieRow, error)
	GetPlansByMovieAndStartTimeOrderByPrice(ctx context.Context, arg *GetPlansByMovieAndStartTimeOrderByPriceParams) ([]*GetPlansByMovieAndStartTimeOrderByPriceRow, error)
	GetPlansCountByTimeWithLock(ctx context.Context, arg *GetPlansCountByTimeWithLockParams) (bool, error)
	GetSeatsByCinemas(ctx context.Context, cinemaID int64) ([]*GetSeatsByCinemasRow, error)
	GetSeatsById(ctx context.Context, id int64) (*GetSeatsByIdRow, error)
	GetTags(ctx context.Context) ([]string, error)
	GetTagsInMovie(ctx context.Context, movieid int64) ([]string, error)
	GetTicket(ctx context.Context, arg *GetTicketParams) (*GetTicketRow, error)
	GetTicketNum(ctx context.Context) (int64, error)
	GetTicketsByPlan(ctx context.Context, planID int64) ([]*GetTicketsByPlanRow, error)
	GetTicketsLocked(ctx context.Context, planid int64) ([]*GetTicketsLockedRow, error)
	GetUserById(ctx context.Context, id int64) (*GetUserByIdRow, error)
	GetUserByName(ctx context.Context, username string) (*GetUserByNameRow, error)
	GetUsers(ctx context.Context) ([]*GetUsersRow, error)
	GetVisitCountsByCreateDate(ctx context.Context, arg *GetVisitCountsByCreateDateParams) (int64, error)
	GetWaitPayOrder(ctx context.Context) ([]*GetWaitPayOrderRow, error)
	ListNameNum(ctx context.Context, username string) (int64, error)
	ListNum(ctx context.Context) (int64, error)
	ListUserInfo(ctx context.Context, arg *ListUserInfoParams) ([]*ListUserInfoRow, error)
	LockTicket(ctx context.Context, arg *LockTicketParams) error
	PayTicket(ctx context.Context, arg *PayTicketParams) error
	QueryCountTicketPlan(ctx context.Context, planID int64) (int64, error)
	SearchAllOrder(ctx context.Context, arg *SearchAllOrderParams) ([]*SearchAllOrderRow, error)
	SearchOrderByCondition(ctx context.Context, arg *SearchOrderByConditionParams) ([]*SearchOrderByConditionRow, error)
	SearchTicketByPlanId(ctx context.Context, arg *SearchTicketByPlanIdParams) ([]*SearchTicketByPlanIdRow, error)
	SearchUserByName(ctx context.Context, arg *SearchUserByNameParams) ([]*SearchUserByNameRow, error)
	UnLockTicket(ctx context.Context, arg *UnLockTicketParams) error
	UpdateCinema(ctx context.Context, arg *UpdateCinemaParams) (*Cinema, error)
	UpdateMovie(ctx context.Context, arg *UpdateMovieParams) (*UpdateMovieRow, error)
	UpdateMovieTag(ctx context.Context, arg *UpdateMovieTagParams) error
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID) error
	UpdatePassword(ctx context.Context, arg *UpdatePasswordParams) error
	UpdateSeats(ctx context.Context, arg *UpdateSeatsParams) error
	UpdateSeatsById(ctx context.Context, arg *UpdateSeatsByIdParams) error
	UpdateUser(ctx context.Context, arg *UpdateUserParams) error
	UpdateUserAvatar(ctx context.Context, arg *UpdateUserAvatarParams) error
}

var _ Querier = (*Queries)(nil)