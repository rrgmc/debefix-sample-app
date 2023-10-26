package payload

type Error struct {
	ErrorMessage string `json:"error_message"`
}

type ErrorValidation struct {
	Error
	Validation map[string]string `json:"validation"`
}
