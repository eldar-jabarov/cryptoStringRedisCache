package dto

type Request struct {
	// example: любая строка
	Query string `json:"query"`
}

type Response struct {
	MD5    string `json:"MD5"`
	SHA256 string `json:"SHA256"`
}
