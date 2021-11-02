package entities

type Publisher struct {
	ID   uint   `json:"-"`
	Name string `json:"name"`
}
