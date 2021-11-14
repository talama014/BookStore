package entities

type Genre struct {
	ID          uint    `json:""`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Book        []*Book `gorm:"many2many:book_genre;"`
}
