package model

type View struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string
	RefKey      string
}

type Views struct {
	Views []View `json:"views"`
}
