package request

type GetVisitCountsByCreateDate struct {
	StartTime int32 `json:"start_time" form:"start_time" binding:"required,gte=1"`
	EndTime   int32 `json:"end_time" form:"end_time" binding:"required,gte=1"`
}
