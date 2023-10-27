package payload

// data

type Country struct {
	CountryID string `json:"country_id"`
	Name      string `json:"name"`
}

// helpers

type CountryFilter struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// request/response

type CountryGetListRequest struct {
	Filter CountryFilter
}

func NewCountryGetListRequest() CountryGetListRequest {
	return CountryGetListRequest{
		Filter: CountryFilter{
			Offset: 0,
			Limit:  100,
		},
	}
}
