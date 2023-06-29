package models

type RequestRows struct {
	Urls []string `json:"urls"`
}

type ResponseRow struct {
	Url          string `json:"url"`
	StatusCode   int    `json:"statuscode"`
	ResponseTime int64  `json:"responsetime"`
}

type ResponseRows struct {
	Urls []ResponseRow `json:"urls"`
}