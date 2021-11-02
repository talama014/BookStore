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

	PublishersID uint      `json:"publisher_id" grom:"-"`
	Publishers   Publisher `json:"publisher,omitempty" gorm:"foreignKey:PublishersID;references:ID"`

	AuthorID uint   `json:"author_id" grom:"-"`
	Authors  Author `json:"authors,omitempty" gorm:"foreignKey:AuthorID;references:ID"`
}
