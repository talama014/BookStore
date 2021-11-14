package entities

import "time"

type Book struct {
	ID             uint      `json:"-"`
	Title          string    `json:"title"`
	Total_pages    int       `json:"total_pages"`
	Isbn           string    `json:"isbn"`
	Quarantine     uint      `json:"quarantine"`
	Publisher_date time.Time `json:"publisher_date"`
	Description    string    `json:"description"`

	PublisherID uint
	Publishers  Publisher `json:"publisher,omitempty" gorm:"foreignKey:PublisherID;references:ID"`

	AuthorID uint
	Authors  Author `json:"author,omitempty" gorm:"foreignKey:AuthorID;references:ID"`
}
