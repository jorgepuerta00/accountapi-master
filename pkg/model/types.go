package model

type PageParams struct {
	Page int
	Size int
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
