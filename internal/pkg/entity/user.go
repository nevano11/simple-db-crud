package entity

type User struct {
	Id       string `json:"-"`
	Name     string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
