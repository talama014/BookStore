package entities

type Author struct {
	ID          uint   `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
