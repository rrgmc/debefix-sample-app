package payload

// data

type Tag struct {
	TagID     string `json:"tag_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TagAdd struct {
	Name string `json:"name"`
}

type TagUpdate = TagAdd

// helpers

type TagFilter struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// request/response

type TagGetListRequest struct {
	Filter TagFilter
}
