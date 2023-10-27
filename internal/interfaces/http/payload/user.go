package payload

// data

type User struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CountryID string `json:"country_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserAdd struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CountryID string `json:"country_id"`
}

type UserUpdate = UserAdd

// helpers

type UserFilter struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// request/response

type UserGetListRequest struct {
	Filter UserFilter
}

func NewUserGetListRequest() UserGetListRequest {
	return UserGetListRequest{
		Filter: UserFilter{
			Offset: 0,
			Limit:  100,
		},
	}
}

type UserAddRequest struct {
	User UserAdd
}

type UserUpdateRequest struct {
	User UserUpdate
}
