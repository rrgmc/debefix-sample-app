package payload

// data

type Comment struct {
	CommentID string `json:"comment_id"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentAdd struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

type CommentUpdate = CommentAdd

// helpers

type CommentFilter struct {
	Offset int     `form:"offset" json:"offset"`
	Limit  int     `form:"limit" json:"limit"`
	PostID *string `form:"post_id" json:"post_id"`
	UserID *string `form:"user_id" json:"user_id"`
}

// request/response

type CommentGetListRequest struct {
	Filter CommentFilter
}

func NewCommentGetListRequest() CommentGetListRequest {
	return CommentGetListRequest{
		Filter: CommentFilter{
			Offset: 0,
			Limit:  100,
		},
	}
}

type CommentAddRequest struct {
	Comment CommentAdd
}

type CommentUpdateRequest struct {
	Comment CommentUpdate
}
