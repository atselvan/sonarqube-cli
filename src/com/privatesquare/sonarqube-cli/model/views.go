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

type ViewComponents struct {
	Results []struct {
		Key      string `json:"key"`
		Selected bool   `json:"selected"`
		Name     string `json:"name"`
	} `json:"results"`
}
