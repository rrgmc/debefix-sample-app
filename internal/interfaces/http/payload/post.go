package payload

// data

type Post struct {
	PostID    string `json:"post_id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	UserID    string `json:"user_id"`
	Tags      []Tag  `json:"tags"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostAdd struct {
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	UserID string   `json:"user_id"`
	Tags   []string `json:"tags"`
}

type PostUpdate = PostAdd

// helpers

type PostFilter struct {
	Offset int     `form:"offset" json:"offset"`
	Limit  int     `form:"limit" json:"limit"`
	UserID *string `form:"user_id" json:"user_id"`
}

// request/response

type PostGetListRequest struct {
	Filter PostFilter
}

func NewPostGetListRequest() PostGetListRequest {
	return PostGetListRequest{
		Filter: PostFilter{
			Offset: 0,
			Limit:  100,
		},
	}
}

type PostAddRequest struct {
	Post PostAdd
}

type PostUpdateRequest struct {
	Post PostUpdate
}
