package request

import "ttms/internal/model/common"

type GetOrderByUserID struct {
	UserID int64 `json:"user_id,omitempty" form:"user_id"`
}

type SearchOrderList struct {
	common.Pager
}

type SearchOrderByCondition struct {
	common.Pager `json:"common_pager" form:"common_pager"`
	Condition    string `json:"condition" form:"condition"`
}
