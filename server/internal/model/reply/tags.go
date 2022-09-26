package reply

type GetTags struct {
	TagsName []string `json:"tags_names,omitempty"`
}

type GetTagsInMovie struct {
	TagName []string `json:"tag_names,omitempty"`
}
