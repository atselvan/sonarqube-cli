package model

type AuthUser struct {
	Username string
	Password string
	Valid    bool `json:"valid"`
}

type UsersAPIResp struct {
	Paging Paging `json:"paging"`
	Users []UserDetails `json:"users"`
}

type UserDetails struct{
	Login            string   `json:"login"`
	Name             string   `json:"name"`
	Active           bool     `json:"active"`
	Email            string   `json:"email"`
	Groups           []string `json:"groups"`
	TokensCount      int      `json:"tokensCount"`
	Local            bool     `json:"local"`
	ExternalIdentity string   `json:"externalIdentity"`
	ExternalProvider string   `json:"externalProvider"`
	Avatar           string   `json:"avatar,omitempty"`
}
