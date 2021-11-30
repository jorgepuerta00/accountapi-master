package model

type PageParams struct {
	page int
	size int
}

type Result struct {
	Data Account `json:"data"`
}

type ArrayResult struct {
	Data []Account `json:"data"`
}

type ErrorResult struct {
	ErrorMessage string `json:"error_message"`
}
