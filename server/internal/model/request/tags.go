package request

type GetTagsInMovie struct {
	MovieID int64 `json:"MovieId" binding:"required,gte=1"`
}

type DeleteOneByMovieAndTag struct {
	MovieID  int64    `json:"MovieId" binding:"required,gte=1"`
	TagNames []string `json:"TagNames" binding:"required,gte=1"`
}

type GetMovieInTag struct {
	TagName string `json:"TagName" binding:"required,gte=1"`
}

type DeleteTagByMovieId struct {
	MovieId int64 `json:"MovieId" binding:"required,gte=1"`
}

type AddNewTagsToMovie struct {
	MovieId  int64    `json:"MovieId" binding:"required,gte=1"`
	TagNames []string `json:"TagNames"`
}

type DeleteByTagName struct {
	TagName string `json:"TagName"`
}

type UpdateMovieTag struct {
	OldTag  string `json:"OldTag"`
	NewTag  string `json:"NewTag"`
	MovieId int64  `json:"MovieId" binding:"required,gte=1"`
}
