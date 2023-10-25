package payload

// data

type Tag struct {
	TagID     string `json:"tag_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TagChange struct {
	Name string `json:"name"`
}

// helpers

type TagFilter struct {
	Offset int
	Limit  int
}

// request/response
