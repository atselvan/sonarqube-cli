package model

type Project struct {
	Id    int    `json:"id"`
	Key   string `json:"k"`
	Name  string `json:"nm"`
	Scope string `json:"sc"`
}
