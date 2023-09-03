package entity

type Car struct {
	Id         int    `json:"-"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Generation int    `json:"generation"`
	Price      int    `json:"price"`
}
