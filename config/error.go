package config

type Error struct {
	Errors []Errors `json:"errors"`
}

type Errors struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Title  string `json:"title"`
}
